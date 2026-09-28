package brew

import (
	"context"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
	"github.com/Automaat/environment-as-code/internal/runner"
	"github.com/Automaat/environment-as-code/internal/runner/runnertest"
)

const brewfile = `# comment
tap "automaat/tap"
brew "jq"
brew "automaat/tap/cache-buster"
cask "ghostty"
mas "Xcode", id: 497799835
`

const cleanupOut = `Would uninstall casks:
zoom
Would uninstall formulae:
wget
Would untap:
old/tap
Would ` + "`brew cleanup`" + `:
Would remove: /Users/x/Library/Caches/Homebrew/foo (1MB)
Run ` + "`brew bundle cleanup --force`" + ` to make these changes.
`

func newModule(t *testing.T, b config.Brew) (*Module, *runnertest.Fake, string) {
	t.Helper()
	root := t.TempDir()
	file := filepath.Join(root, "Brewfile")
	if err := os.WriteFile(file, []byte(brewfile), 0o644); err != nil {
		t.Fatal(err)
	}
	b.File = "Brewfile"
	fake := runnertest.New()
	return &Module{Brew: b, Paths: config.Paths{Root: root}, Runner: fake}, fake, file
}

func targets(changes []engine.Change) []string {
	var out []string
	for _, c := range changes {
		out = append(out, string(c.Action)+" "+c.Target)
	}
	return out
}

func TestParseBrewfile(t *testing.T) {
	_, _, file := newModule(t, config.Brew{})
	got, err := ParseBrewfile(file)
	if err != nil {
		t.Fatal(err)
	}
	want := []Entry{{"tap", "automaat/tap"}, {"brew", "jq"}, {"brew", "automaat/tap/cache-buster"}, {"cask", "ghostty"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
}

func TestParseCleanup(t *testing.T) {
	got := ParseCleanup(cleanupOut)
	want := []string{"cask zoom", "brew wget", "tap old/tap"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	if got := ParseCleanup(""); got != nil {
		t.Errorf("empty output: %v", got)
	}
}

func TestPlanUpToDate(t *testing.T) {
	m, fake, file := newModule(t, config.Brew{Cleanup: config.CleanupZap})
	fake.OnOK("brew trust --json=v1", `{"taps":["automaat/tap"],"formulae":[]}`)
	fake.OnOK("brew bundle check --file "+file+" --verbose --no-upgrade", "The Brewfile's dependencies are satisfied.\n")
	fake.OnOK("brew bundle cleanup --file "+file, "")

	changes, err := m.Plan(context.Background())
	if err != nil || len(changes) != 0 {
		t.Fatalf("changes = %v, err = %v", changes, err)
	}
	for _, c := range fake.Calls {
		if !reflect.DeepEqual(c.Env, noAutoUpdate) {
			t.Errorf("%s ran without disabling auto-update", c)
		}
	}
}

func TestPlanInstallUpgradeAndZap(t *testing.T) {
	m, fake, file := newModule(t, config.Brew{Cleanup: config.CleanupZap, Upgrade: true})
	fake.OnOK("brew trust --json=v1", `{"taps":[],"formulae":["automaat/tap/cache-buster"]}`)
	fake.OnOK("brew trust --tap automaat/tap", "")
	fake.On("brew bundle check --file "+file+" --verbose --no-upgrade", runner.Result{ExitCode: 1, Stdout: `brew bundle can't satisfy your Brewfile's dependencies.
→ Cask ghostty needs to be installed.
→ Formula jq needs to be installed or updated.
→ Tap new/tap needs to be tapped.
Satisfy missing dependencies with ` + "`brew bundle install`."})
	fake.OnOK("brew outdated --json=v2", `{"formulae":[{"name":"cache-buster"},{"name":"not-in-brewfile"}],"casks":[{"name":"ghostty"}]}`)
	fake.OnOK("brew bundle cleanup --file "+file, cleanupOut)
	fake.OnOK("brew bundle install --file "+file, "")
	fake.OnOK("brew bundle cleanup --force --file "+file+" --zap", "")

	changes, err := m.Plan(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"+ trust tap automaat/tap",
		"+ cask ghostty", "+ formula jq", "+ tap new/tap",
		"~ brew cache-buster", "~ cask ghostty",
		"! brew bundle install",
		"- cask zoom", "- brew wget", "- tap old/tap",
		"! brew bundle cleanup",
	}
	if got := targets(changes); !reflect.DeepEqual(got, want) {
		t.Fatalf("got  %v\nwant %v", got, want)
	}

	if err := engine.Apply(context.Background(), func(string) {}, engine.Plan{{Module: "brew", Changes: changes}}); err != nil {
		t.Fatal(err)
	}
	lines := fake.Lines()
	if got := lines[len(lines)-3:]; !reflect.DeepEqual(got, []string{
		"brew trust --tap automaat/tap",
		"brew bundle install --file " + file,
		"brew bundle cleanup --force --file " + file + " --zap",
	}) {
		t.Errorf("applied %v", got)
	}
}

func TestNoUpgradeAndNoCleanup(t *testing.T) {
	m, fake, file := newModule(t, config.Brew{Cleanup: config.CleanupNone})
	fake.OnOK("brew trust --json=v1", `{"taps":["automaat/tap"]}`)
	fake.On("brew bundle check --file "+file+" --verbose --no-upgrade", runner.Result{ExitCode: 1, Stdout: "→ Formula jq needs to be installed.\n"})
	fake.OnOK("brew bundle install --file "+file+" --no-upgrade", "")

	changes, err := m.Plan(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if got := targets(changes); !reflect.DeepEqual(got, []string{"+ formula jq", "! brew bundle install"}) {
		t.Fatalf("got %v", got)
	}
	if err := engine.Apply(context.Background(), func(string) {}, engine.Plan{{Changes: changes}}); err != nil {
		t.Fatal(err)
	}
	if fake.Ran("brew outdated") || fake.Ran("brew bundle cleanup") {
		t.Errorf("unexpected calls: %v", fake.Lines())
	}
}

func TestCheckFailureWithoutMissingEntries(t *testing.T) {
	m, fake, file := newModule(t, config.Brew{})
	fake.OnOK("brew trust --json=v1", `{"taps":["automaat/tap"]}`)
	fake.On("brew bundle check --file "+file+" --verbose --no-upgrade", runner.Result{ExitCode: 1, Stderr: "Error: invalid Brewfile"})

	_, err := m.Plan(context.Background())
	if err == nil || !strings.Contains(err.Error(), "invalid Brewfile") {
		t.Errorf("err = %v", err)
	}
}
