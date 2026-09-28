package mise

import (
	"context"
	"reflect"
	"testing"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
	"github.com/Automaat/environment-as-code/internal/runner"
	"github.com/Automaat/environment-as-code/internal/runner/runnertest"
)

func newModule() (*Module, *runnertest.Fake) {
	fake := runnertest.New()
	return &Module{
		Mise:   config.Mise{Config: "dotfiles/mise/config.toml"},
		Paths:  config.Paths{Home: "/home/u", Root: "/repo"},
		Runner: fake,
	}, fake
}

func TestPlanAndApply(t *testing.T) {
	m, fake := newModule()
	fake.OnOK("mise ls --global --missing --json", `{
  "kubectl": [{"version": "1.34.1", "installed": false}],
  "jq": [{"version": "1.7.1", "installed": false}, {"version": "1.6", "installed": true}]
}`)
	fake.OnOK("mise install --yes", "")

	changes, err := m.Plan(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, c := range changes {
		got = append(got, string(c.Action)+" "+c.Target)
	}
	want := []string{"+ jq@1.7.1", "+ kubectl@1.34.1", "! mise install"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	if err := engine.Apply(context.Background(), func(string) {}, engine.Plan{{Changes: changes}}); err != nil {
		t.Fatal(err)
	}
	for _, c := range fake.Calls {
		if c.Dir != "/home/u" || !reflect.DeepEqual(c.Env, []string{"MISE_GLOBAL_CONFIG_FILE=/repo/dotfiles/mise/config.toml"}) {
			t.Errorf("%s: dir %q env %v; must target the repo config from home", c, c.Dir, c.Env)
		}
	}
	if !fake.Calls[len(fake.Calls)-1].Stream {
		t.Error("install should stream output")
	}
}

func TestNothingMissing(t *testing.T) {
	m, fake := newModule()
	fake.OnOK("mise ls --global --missing --json", "{}")
	changes, err := m.Plan(context.Background())
	if err != nil || len(changes) != 0 {
		t.Errorf("changes = %v, err = %v", changes, err)
	}
}

func TestPlanErrors(t *testing.T) {
	tests := map[string]runner.Result{
		"command fails": {ExitCode: 1, Stderr: "config not trusted"},
		"bad json":      {Stdout: "not json"},
	}
	for name, res := range tests {
		t.Run(name, func(t *testing.T) {
			m, fake := newModule()
			fake.On("mise ls --global --missing --json", res)
			if _, err := m.Plan(context.Background()); err == nil {
				t.Error("expected error")
			}
		})
	}
}
