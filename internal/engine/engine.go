// Package engine plans and applies changes across modules.
package engine

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Action classifies a change for display.
type Action string

const (
	Create Action = "+"
	Update Action = "~"
	Remove Action = "-"
	Run    Action = "!"
)

// Change is one difference between desired and actual state. Apply is nil for
// changes that are only informational and applied by a later Run change of
// the same module (e.g. every missing brew entry is installed by one
// `brew bundle install`).
type Change struct {
	Action Action
	Target string
	Detail string
	Apply  func(ctx context.Context) error
}

func (c Change) String() string {
	if c.Detail == "" {
		return fmt.Sprintf("%s %s", c.Action, c.Target)
	}
	return fmt.Sprintf("%s %s (%s)", c.Action, c.Target, c.Detail)
}

// Module computes the changes needed to converge one area of the system.
type Module interface {
	Name() string
	Plan(ctx context.Context) ([]Change, error)
}

// ModulePlan is the plan for one module.
type ModulePlan struct {
	Module  string
	Changes []Change
}

// Plan is the full set of pending changes, in module order.
type Plan []ModulePlan

// Empty reports whether nothing needs to change.
func (p Plan) Empty() bool {
	return p.Count() == 0
}

// Count returns the number of changes across all modules.
func (p Plan) Count() int {
	n := 0
	for _, mp := range p {
		n += len(mp.Changes)
	}
	return n
}

// Build plans every module in order, stopping at the first failure.
func Build(ctx context.Context, modules []Module) (Plan, error) {
	plan := make(Plan, 0, len(modules))
	for _, m := range modules {
		changes, err := m.Plan(ctx)
		if err != nil {
			return nil, fmt.Errorf("plan %s: %w", m.Name(), err)
		}
		plan = append(plan, ModulePlan{Module: m.Name(), Changes: changes})
	}
	return plan, nil
}

// Print writes a human-readable plan.
func Print(w io.Writer, p Plan) error {
	var b strings.Builder
	for _, mp := range p {
		if len(mp.Changes) == 0 {
			fmt.Fprintf(&b, "%s: up to date\n", mp.Module)
			continue
		}
		fmt.Fprintf(&b, "%s:\n", mp.Module)
		for _, c := range mp.Changes {
			fmt.Fprintf(&b, "  %s\n", c)
		}
	}
	_, err := io.WriteString(w, b.String())
	return err
}

// Apply executes the plan in order, reporting each change to progress before
// running it. Within a module the first failure stops that module, since
// later changes often depend on earlier ones; other modules still run and all
// failures are returned together.
func Apply(ctx context.Context, progress func(string), p Plan) error {
	var errs []error
	for _, mp := range p {
		for _, c := range mp.Changes {
			if c.Apply == nil {
				continue
			}
			progress(mp.Module + ": " + c.String())
			if err := c.Apply(ctx); err != nil {
				errs = append(errs, fmt.Errorf("%s: %s: %w", mp.Module, c.Target, err))
				break
			}
		}
	}
	return errors.Join(errs...)
}
