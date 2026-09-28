// Package mise installs the pinned tool versions from the global mise config.
package mise

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
	"github.com/Automaat/environment-as-code/internal/runner"
)

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

// The repo config is pointed at explicitly, and commands run from home, so
// the plan does not depend on the link being in place yet or on whichever
// project config the caller's cwd would pick up.
func (m *Module) cmd(args ...string) runner.Cmd {
	return runner.Cmd{
		Name: "mise",
		Args: args,
		Dir:  m.Paths.Home,
		Env:  []string{"MISE_GLOBAL_CONFIG_FILE=" + m.Paths.Src(m.Mise.Config)},
	}
}

func (m *Module) Plan(ctx context.Context) ([]engine.Change, error) {
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
