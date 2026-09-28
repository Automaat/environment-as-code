// Package mise installs the pinned tool versions from the global mise config
// and optionally prunes versions no config references anymore.
package mise

import (
	"context"
	"encoding/json"
	"fmt"
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
	prune, err := m.planPrune(ctx)
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

func (m *Module) planPrune(ctx context.Context) ([]engine.Change, error) {
	res, err := m.Runner.Run(ctx, m.cmd("prune", "--dry-run"))
	if err != nil {
		return nil, err
	}
	if err := res.Err(m.cmd("prune", "--dry-run")); err != nil {
		return nil, err
	}
	prunable := ParsePrunable(res.Stdout + res.Stderr)
	if len(prunable) == 0 {
		return nil, nil
	}
	var changes []engine.Change
	for _, p := range prunable {
		changes = append(changes, engine.Change{Action: engine.Remove, Target: p})
	}
	prune := m.cmd("prune", "--yes")
	prune.Stream = true
	return append(changes, engine.Change{
		Action: engine.Run, Target: "mise prune",
		Apply: func(ctx context.Context) error { return runner.Check(ctx, m.Runner, prune) },
	}), nil
}
