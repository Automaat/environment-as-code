// Package mise installs the pinned tool versions from the global mise config
// and optionally prunes versions no config references anymore.
package mise

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
	"github.com/Automaat/environment-as-code/internal/runner"
)

// Dir is where mise commands run. From $HOME, mise treats
// ~/.config/mise/config.toml as a project config that outranks the repo file
// passed in MISE_GLOBAL_CONFIG_FILE, so bumped pins would look installed
// until the files module had copied the new config there.
const Dir = "/"

type Module struct {
	Mise   config.Mise
	Paths  config.Paths
	Runner runner.Runner
}

func (m *Module) Name() string { return "mise" }

type tool struct {
	Version   string `json:"version"`
	Installed bool   `json:"installed"`
}

func (m *Module) cmd(args ...string) runner.Cmd {
	return runner.Cmd{
		Name: "mise",
		Args: args,
		Dir:  Dir,
		Env:  []string{"MISE_GLOBAL_CONFIG_FILE=" + m.Paths.Src(m.Mise.Config)},
	}
}

func (m *Module) Plan(ctx context.Context) ([]engine.Change, error) {
	changes, err := m.planInstall(ctx)
	if err != nil {
		return nil, err
	}
	if !m.Mise.Prune {
		return changes, nil
	}
	prune, err := m.planPrune(ctx, len(changes) > 0)
	if err != nil {
		return nil, err
	}
	return append(changes, prune...), nil
}

func (m *Module) planInstall(ctx context.Context) ([]engine.Change, error) {
	out, err := runner.Output(ctx, m.Runner, m.cmd("ls", "--global", "--missing", "--json"))
	if err != nil {
		return nil, err
	}
	var tools map[string][]tool
	if err := json.Unmarshal([]byte(out), &tools); err != nil {
		return nil, fmt.Errorf("mise ls: %w", err)
	}

	names := make([]string, 0, len(tools))
	for name := range tools {
		names = append(names, name)
	}
	sort.Strings(names)

	var changes []engine.Change
	for _, name := range names {
		for _, t := range tools[name] {
			if !t.Installed {
				changes = append(changes, engine.Change{Action: engine.Create, Target: name + "@" + t.Version})
			}
		}
	}
	if len(changes) == 0 {
		return nil, nil
	}
	install := m.cmd("install", "--yes")
	install.Stream = true
	return append(changes, engine.Change{
		Action: engine.Run, Target: "mise install",
		Apply: func(ctx context.Context) error { return runner.Check(ctx, m.Runner, install) },
	}), nil
}

var prunableLine = regexp.MustCompile(`^(?:mise\s+)?(\S+@\S+) is prunable`)

// ParsePrunable extracts tool@version entries from `mise prune --dry-run`.
func ParsePrunable(out string) []string {
	var res []string
	for line := range strings.Lines(out) {
		if g := prunableLine.FindStringSubmatch(strings.TrimSpace(line)); g != nil {
			res = append(res, g[1])
		}
	}
	return res
}

// planPrune asks mise what it would prune now, but mise also counts the
// installed ~/.config/mise/config.toml, which still holds the old pins until
// the files module replaces it. So when tools are being installed or that
// copy is stale, prune is scheduled anyway: at apply time it runs after the
// new config is in place and removes the superseded versions in one pass.
func (m *Module) planPrune(ctx context.Context, installing bool) ([]engine.Change, error) {
	dryRun := m.cmd("prune", "--dry-run")
	res, err := m.Runner.Run(ctx, dryRun)
	if err != nil {
		return nil, err
	}
	if err := res.Err(dryRun); err != nil {
		return nil, err
	}
	prunable := ParsePrunable(res.Stdout + res.Stderr)
	stale, err := m.installedConfigStale()
	if err != nil {
		return nil, err
	}
	if len(prunable) == 0 && !installing && !stale {
		return nil, nil
	}
	var changes []engine.Change
	for _, p := range prunable {
		changes = append(changes, engine.Change{Action: engine.Remove, Target: p})
	}
	detail := ""
	if len(prunable) == 0 {
		detail = "versions superseded by the config update"
	}
	prune := m.cmd("prune", "--yes")
	prune.Stream = true
	return append(changes, engine.Change{
		Action: engine.Run, Target: "mise prune", Detail: detail,
		Apply: func(ctx context.Context) error { return runner.Check(ctx, m.Runner, prune) },
	}), nil
}

// InstalledConfig is where mise reads the user's global config; the files
// module keeps it in sync with the repo copy.
const InstalledConfig = "~/.config/mise/config.toml"

func (m *Module) installedConfigStale() (bool, error) {
	want, err := os.ReadFile(m.Paths.Src(m.Mise.Config))
	if err != nil {
		return false, err
	}
	have, err := os.ReadFile(m.Paths.Dst(InstalledConfig))
	if errors.Is(err, fs.ErrNotExist) {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return !bytes.Equal(want, have), nil
}
