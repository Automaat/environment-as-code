// Package commands runs guarded shell snippets: `run` executes only while
// `check` fails. It is the escape hatch for setup no other module covers.
package commands

import (
	"context"
	"fmt"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
	"github.com/Automaat/environment-as-code/internal/runner"
)

type Module struct {
	Commands []config.Command
	Runner   runner.Runner
}

func (m *Module) Name() string { return "commands" }

func (m *Module) Plan(ctx context.Context) ([]engine.Change, error) {
	var changes []engine.Change
	for _, c := range m.Commands {
		ok, err := m.satisfied(ctx, c)
		if err != nil {
			return nil, err
		}
		if ok {
			continue
		}
		changes = append(changes, engine.Change{Action: engine.Run, Target: c.Name, Detail: c.Run, Apply: m.apply(c)})
	}
	return changes, nil
}

func (m *Module) satisfied(ctx context.Context, c config.Command) (bool, error) {
	res, err := m.Runner.Run(ctx, runner.Cmd{Name: "sh", Args: []string{"-c", c.Check}})
	if err != nil {
		return false, err
	}
	return res.ExitCode == 0, nil
}

// apply re-checks first: an earlier module (brew, mise) may have satisfied
// the check since planning.
func (m *Module) apply(c config.Command) func(context.Context) error {
	return func(ctx context.Context) error {
		ok, err := m.satisfied(ctx, c)
		if err != nil || ok {
			return err
		}
		if err := runner.Check(ctx, m.Runner, runner.Cmd{Name: "sh", Args: []string{"-c", c.Run}, Stream: true}); err != nil {
			return err
		}
		if ok, err := m.satisfied(ctx, c); err != nil {
			return err
		} else if !ok {
			return fmt.Errorf("check %q still fails after run", c.Check)
		}
		return nil
	}
}
