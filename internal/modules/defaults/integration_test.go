//go:build integration

package defaults

import (
	"context"
	"os/exec"
	"testing"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
	"github.com/Automaat/environment-as-code/internal/runner"
)

const domain = "dev.eac.integration-test"

// TestRealDefaults round-trips every supported type through the real
// `defaults` CLI, proving Encode matches what `defaults read` prints.
func TestRealDefaults(t *testing.T) {
	if _, err := exec.LookPath("defaults"); err != nil {
		t.Skip("defaults CLI not available")
	}
	_ = exec.Command("defaults", "delete", domain).Run()
	t.Cleanup(func() { _ = exec.Command("defaults", "delete", domain).Run() })

	m := &Module{Runner: runner.NewExec(), Defaults: []config.Default{
		{Domain: domain, Key: "b", Value: true},
		{Domain: domain, Key: "i", Value: 15},
		{Domain: domain, Key: "f", Value: 0.25},
		{Domain: domain, Key: "s", Value: "~/Documents/screenshots"},
	}}
	ctx := context.Background()

	changes, err := m.Plan(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(changes) != 4 {
		t.Fatalf("fresh domain: %v", changes)
	}
	if err := engine.Apply(ctx, func(string) {}, engine.Plan{{Changes: changes}}); err != nil {
		t.Fatal(err)
	}
	if again, err := m.Plan(ctx); err != nil || len(again) != 0 {
		t.Fatalf("not idempotent: %v, %v", again, err)
	}

	if err := exec.Command("defaults", "write", domain, "i", "-string", "15").Run(); err != nil {
		t.Fatal(err)
	}
	drift, err := m.Plan(ctx)
	if err != nil || len(drift) != 1 || drift[0].Detail != "string:15 → int:15" {
		t.Errorf("type drift not detected: %v, %v", drift, err)
	}
}
