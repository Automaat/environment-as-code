// Package brew converges Homebrew to a Brewfile via `brew bundle`.
package brew

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"regexp"
	"slices"
	"strings"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
	"github.com/Automaat/environment-as-code/internal/runner"
)

// Planning must be read-only and fast, so brew never self-updates while
// we only inspect state.
var noAutoUpdate = []string{"HOMEBREW_NO_AUTO_UPDATE=1", "HOMEBREW_NO_ENV_HINTS=1"}

type Module struct {
	Brew   config.Brew
	Paths  config.Paths
	Runner runner.Runner
}

func (m *Module) Name() string { return "brew" }

func (m *Module) Plan(ctx context.Context) ([]engine.Change, error) {
	file := m.Paths.Src(m.Brew.File)
	entries, err := ParseBrewfile(file)
	if err != nil {
		return nil, err
	}

	trust, err := m.planTrust(ctx, entries)
	if err != nil {
		return nil, err
	}

	var install []engine.Change
	missing, err := m.missing(ctx, file)
	if err != nil {
		return nil, err
	}
	for _, e := range missing {
		install = append(install, engine.Change{Action: engine.Create, Target: e})
	}
	if m.Brew.Upgrade {
		outdated, err := m.outdated(ctx, entries)
		if err != nil {
			return nil, err
		}
		for _, e := range outdated {
			install = append(install, engine.Change{Action: engine.Update, Target: e, Detail: "outdated"})
		}
	}
	if len(install) > 0 {
		install = append(install, m.bundleInstall(file))
	}

	cleanup, err := m.planCleanup(ctx, file)
	if err != nil {
		return nil, err
	}
	return append(append(trust, install...), cleanup...), nil
}

type trustJSON struct {
	Taps []string `json:"taps"`
}

// planTrust trusts every tap the Brewfile lists: Homebrew refuses to load
// formulae from untrusted third-party taps, and listing a tap in the
// Brewfile already states the intent to use it.
func (m *Module) planTrust(ctx context.Context, entries []Entry) ([]engine.Change, error) {
	var taps []string
	for _, e := range entries {
		if e.Kind == "tap" && !strings.HasPrefix(e.Name, "homebrew/") {
			taps = append(taps, e.Name)
		}
	}
	if len(taps) == 0 {
		return nil, nil
	}
	out, err := runner.Output(ctx, m.Runner, runner.Cmd{Name: "brew", Args: []string{"trust", "--json=v1"}, Env: noAutoUpdate})
	if err != nil {
		return nil, err
	}
	var trusted trustJSON
	if err := json.Unmarshal([]byte(out), &trusted); err != nil {
		return nil, fmt.Errorf("brew trust: %w", err)
	}
	var changes []engine.Change
	for _, tap := range taps {
		if slices.Contains(trusted.Taps, tap) {
			continue
		}
		cmd := runner.Cmd{Name: "brew", Args: []string{"trust", "--tap", tap}, Env: noAutoUpdate}
		changes = append(changes, engine.Change{
			Action: engine.Create, Target: "trust tap " + tap,
			Apply: func(ctx context.Context) error { return runner.Check(ctx, m.Runner, cmd) },
		})
	}
	return changes, nil
}

func (m *Module) bundleInstall(file string) engine.Change {
	args := []string{"bundle", "install", "--file", file}
	if !m.Brew.Upgrade {
		args = append(args, "--no-upgrade")
	}
	cmd := runner.Cmd{Name: "brew", Args: args, Stream: true}
	return engine.Change{
		Action: engine.Run, Target: "brew bundle install",
		Apply: func(ctx context.Context) error { return runner.Check(ctx, m.Runner, cmd) },
	}
}

var missingLine = regexp.MustCompile(`^→ (\S+) (\S+) needs to be `)

func (m *Module) missing(ctx context.Context, file string) ([]string, error) {
	cmd := runner.Cmd{Name: "brew", Args: []string{"bundle", "check", "--file", file, "--verbose", "--no-upgrade"}, Env: noAutoUpdate}
	res, err := m.Runner.Run(ctx, cmd)
	if err != nil {
		return nil, err
	}
	if res.ExitCode == 0 {
		return nil, nil
	}
	var out []string
	for line := range strings.Lines(res.Stdout + res.Stderr) {
		if g := missingLine.FindStringSubmatch(strings.TrimSpace(line)); g != nil {
			out = append(out, strings.ToLower(g[1])+" "+g[2])
		}
	}
	if len(out) == 0 {
		return nil, res.Err(cmd)
	}
	return out, nil
}

type outdatedJSON struct {
	Formulae []struct {
		Name string `json:"name"`
	} `json:"formulae"`
	Casks []struct {
		Name string `json:"name"`
	} `json:"casks"`
}

func (m *Module) outdated(ctx context.Context, entries []Entry) ([]string, error) {
	out, err := runner.Output(ctx, m.Runner, runner.Cmd{Name: "brew", Args: []string{"outdated", "--json=v2"}, Env: noAutoUpdate})
	if err != nil {
		return nil, err
	}
	var parsed outdatedJSON
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		return nil, fmt.Errorf("brew outdated: %w", err)
	}
	wanted := map[string]bool{}
	for _, e := range entries {
		wanted[e.Kind+" "+path.Base(e.Name)] = true
	}
	var res []string
	for _, f := range parsed.Formulae {
		if key := "brew " + path.Base(f.Name); wanted[key] {
			res = append(res, key)
		}
	}
	for _, c := range parsed.Casks {
		if key := "cask " + path.Base(c.Name); wanted[key] {
			res = append(res, key)
		}
	}
	return res, nil
}

func (m *Module) planCleanup(ctx context.Context, file string) ([]engine.Change, error) {
	mode := m.Brew.Cleanup
	if mode == "" || mode == config.CleanupNone {
		return nil, nil
	}
	out, err := runner.Output(ctx, m.Runner, runner.Cmd{Name: "brew", Args: []string{"bundle", "cleanup", "--file", file}, Env: noAutoUpdate})
	if err != nil {
		return nil, err
	}
	removals := ParseCleanup(out)
	if len(removals) == 0 {
		return nil, nil
	}
	var changes []engine.Change
	for _, r := range removals {
		changes = append(changes, engine.Change{Action: engine.Remove, Target: r})
	}
	args := []string{"bundle", "cleanup", "--force", "--file", file}
	if mode == config.CleanupZap {
		args = append(args, "--zap")
	}
	cmd := runner.Cmd{Name: "brew", Args: args, Stream: true}
	return append(changes, engine.Change{
		Action: engine.Run, Target: "brew bundle cleanup",
		Detail: mode,
		Apply:  func(ctx context.Context) error { return runner.Check(ctx, m.Runner, cmd) },
	}), nil
}

var cleanupSections = map[string]string{
	"Would uninstall casks:":    "cask",
	"Would uninstall formulae:": "brew",
	"Would untap:":              "tap",
}

// ParseCleanup extracts what `brew bundle cleanup` (dry run) would remove.
// Cache pruning ("Would `brew cleanup`") is ignored as it is not state we
// manage.
func ParseCleanup(out string) []string {
	var res []string
	kind := ""
	for line := range strings.Lines(out) {
		line = strings.TrimSpace(line)
		if k, ok := cleanupSections[line]; ok {
			kind = k
			continue
		}
		if strings.HasPrefix(line, "Would ") || strings.HasPrefix(line, "Run ") || line == "" {
			kind = ""
			continue
		}
		if kind != "" {
			res = append(res, kind+" "+line)
		}
	}
	return res
}

// Entry is one tap/brew/cask line of a Brewfile.
type Entry struct {
	Kind string
	Name string
}

var entryLine = regexp.MustCompile(`^(tap|brew|cask)\s+"([^"]+)"`)

// ParseBrewfile reads the tap/brew/cask entries of a Brewfile. Other
// directives (mas, vscode, …) are left to brew bundle.
func ParseBrewfile(file string) ([]Entry, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var entries []Entry
	for line := range strings.Lines(string(data)) {
		if g := entryLine.FindStringSubmatch(strings.TrimSpace(line)); g != nil {
			entries = append(entries, Entry{Kind: g[1], Name: g[2]})
		}
	}
	return entries, nil
}
