package commands

import (
	"context"
	"strings"
	"testing"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
	"github.com/Automaat/environment-as-code/internal/runner"
	"github.com/Automaat/environment-as-code/internal/runner/runnertest"
)

var (
	pass = runner.Result{}
	fail = runner.Result{ExitCode: 1}
)

func cmd() config.Command {
	return config.Command{Name: "gke-auth", Check: "command -v gke-gcloud-auth-plugin", Run: "gcloud components install gke-gcloud-auth-plugin"}
}

func run(t *testing.T, fake *runnertest.Fake) ([]engine.Change, error) {
	t.Helper()
	m := &Module{Runner: fake, Commands: []config.Command{cmd()}}
	changes, err := m.Plan(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return changes, engine.Apply(context.Background(), func(string) {}, engine.Plan{{Changes: changes}})
}

func TestSatisfiedIsSkipped(t *testing.T) {
	fake := runnertest.New().On("sh -c "+cmd().Check, pass)
	changes, err := run(t, fake)
	if err != nil || len(changes) != 0 {
		t.Errorf("changes = %v, err = %v", changes, err)
	}
}

func TestRunsUntilCheckPasses(t *testing.T) {
	fake := runnertest.New().
		On("sh -c "+cmd().Check, fail).
		On("sh -c "+cmd().Check, fail).
		On("sh -c "+cmd().Check, pass).
		On("sh -c "+cmd().Run, pass)

	changes, err := run(t, fake)
	if err != nil || len(changes) != 1 {
		t.Fatalf("changes = %v, err = %v", changes, err)
	}
	if !fake.Ran("sh -c " + cmd().Run) {
		t.Error("run was not executed")
	}
}

func TestSatisfiedSincePlanning(t *testing.T) {
	fake := runnertest.New().
		On("sh -c "+cmd().Check, fail).
		On("sh -c "+cmd().Check, pass)

	if _, err := run(t, fake); err != nil {
		t.Fatal(err)
	}
	if fake.Ran("sh -c " + cmd().Run) {
		t.Error("run must be skipped when an earlier module already satisfied the check")
	}
}

func TestStillFailingAfterRun(t *testing.T) {
	fake := runnertest.New().
		On("sh -c "+cmd().Check, fail).
		On("sh -c "+cmd().Run, pass)

	_, err := run(t, fake)
	if err == nil || !strings.Contains(err.Error(), "still fails") {
		t.Errorf("err = %v", err)
	}
}

func TestRunFailure(t *testing.T) {
	fake := runnertest.New().
		On("sh -c "+cmd().Check, fail).
		On("sh -c "+cmd().Run, runner.Result{ExitCode: 2, Stderr: "network down"})

	_, err := run(t, fake)
	if err == nil || !strings.Contains(err.Error(), "network down") {
		t.Errorf("err = %v", err)
	}
}
