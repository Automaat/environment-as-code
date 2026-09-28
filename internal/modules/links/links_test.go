package links

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
	"github.com/Automaat/environment-as-code/internal/install"
)

type fixture struct {
	home, root string
	mod        *Module
}

func newFixture(t *testing.T) *fixture {
	t.Helper()
	f := &fixture{home: t.TempDir(), root: t.TempDir()}
	f.write(t, filepath.Join(f.root, "dotfiles/zshrc"), "repo")
	f.mod = &Module{
		Links: []config.Link{{Src: "dotfiles/zshrc", Dst: "~/.config/zsh/.zshrc"}},
		Paths: config.Paths{Home: f.home, Root: f.root},
		Now:   func() time.Time { return time.Unix(1700000000, 0) },
	}
	return f
}

func (f *fixture) write(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func (f *fixture) dst() string { return filepath.Join(f.home, ".config/zsh/.zshrc") }
func (f *fixture) src() string { return filepath.Join(f.root, "dotfiles/zshrc") }

func converge(t *testing.T, m engine.Module) []engine.Change {
	t.Helper()
	ctx := context.Background()
	changes, err := m.Plan(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := engine.Apply(ctx, func(string) {}, engine.Plan{{Module: m.Name(), Changes: changes}}); err != nil {
		t.Fatal(err)
	}
	again, err := m.Plan(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(again) != 0 {
		t.Fatalf("not idempotent, second plan: %v", again)
	}
	return changes
}

func TestLinks(t *testing.T) {
	tests := []struct {
		name       string
		setup      func(t *testing.T, f *fixture)
		wantAction engine.Action
		wantDetail string
		check      func(t *testing.T, f *fixture)
	}{
		{
			name:       "missing destination is created with parents",
			setup:      func(*testing.T, *fixture) {},
			wantAction: engine.Create,
			wantDetail: "→ ",
		},
		{
			name: "symlink elsewhere is relinked",
			setup: func(t *testing.T, f *fixture) {
				if err := os.MkdirAll(filepath.Dir(f.dst()), 0o755); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink("/nix/store/abc-home-manager-files/.zshrc", f.dst()); err != nil {
					t.Fatal(err)
				}
			},
			wantAction: engine.Update,
			wantDetail: "relink /nix/store/abc-home-manager-files/.zshrc",
		},
		{
			name:       "regular file is backed up",
			setup:      func(t *testing.T, f *fixture) { f.write(t, f.dst(), "local edits") },
			wantAction: engine.Update,
			wantDetail: "back up to .zshrc.eac-bak",
			check: func(t *testing.T, f *fixture) {
				got, err := os.ReadFile(f.dst() + ".eac-bak")
				if err != nil || string(got) != "local edits" {
					t.Errorf("backup = %q, %v", got, err)
				}
			},
		},
		{
			name: "existing backup gets a timestamped name",
			setup: func(t *testing.T, f *fixture) {
				f.write(t, f.dst(), "new edits")
				f.write(t, f.dst()+".eac-bak", "old backup")
			},
			wantAction: engine.Update,
			wantDetail: "back up to .zshrc.eac-bak.1700000000",
			check: func(t *testing.T, f *fixture) {
				got, _ := os.ReadFile(f.dst() + ".eac-bak.1700000000")
				old, _ := os.ReadFile(f.dst() + ".eac-bak")
				if string(got) != "new edits" || string(old) != "old backup" {
					t.Errorf("backups = %q, %q", got, old)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := newFixture(t)
			tt.setup(t, f)

			changes := converge(t, f.mod)

			if len(changes) != 1 {
				t.Fatalf("changes = %v", changes)
			}
			c := changes[0]
			if c.Action != tt.wantAction || !strings.Contains(c.Detail, tt.wantDetail) || c.Target != "~/.config/zsh/.zshrc" {
				t.Errorf("change = %s", c)
			}
			if target, err := os.Readlink(f.dst()); err != nil || target != f.src() {
				t.Errorf("link = %q, %v", target, err)
			}
			if tt.check != nil {
				tt.check(t, f)
			}
		})
	}
}

func TestLinksErrors(t *testing.T) {
	t.Run("missing source", func(t *testing.T) {
		f := newFixture(t)
		f.mod.Links[0].Src = "dotfiles/nope"
		if _, err := f.mod.Plan(context.Background()); err == nil || !strings.Contains(err.Error(), "source") {
			t.Errorf("err = %v", err)
		}
	})
	t.Run("directory in the way", func(t *testing.T) {
		f := newFixture(t)
		if err := os.MkdirAll(f.dst(), 0o755); err != nil {
			t.Fatal(err)
		}
		if _, err := f.mod.Plan(context.Background()); err == nil || !strings.Contains(err.Error(), "is a directory") {
			t.Errorf("err = %v", err)
		}
	})
}

func TestLinkDirectory(t *testing.T) {
	f := newFixture(t)
	f.write(t, filepath.Join(f.root, "dotfiles/hammerspoon/init.lua"), "lua")
	f.mod.Links = []config.Link{{Src: "dotfiles/hammerspoon", Dst: "~/.hammerspoon"}}

	converge(t, f.mod)

	got, err := os.ReadFile(filepath.Join(f.home, ".hammerspoon/init.lua"))
	if err != nil || string(got) != "lua" {
		t.Errorf("read through dir link = %q, %v", got, err)
	}
}

func TestTakesOverProtectedFile(t *testing.T) {
	f := newFixture(t)
	t.Cleanup(func() { _ = install.Unlock(f.home) })
	state, err := install.LoadState(install.StatePath(f.home))
	if err != nil {
		t.Fatal(err)
	}
	in := &install.Installer{Paths: f.mod.Paths, State: state, Immutable: true}
	c, err := in.Plan(f.dst(), []byte("locked copy"), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Apply(context.Background()); err != nil {
		t.Fatal(err)
	}

	converge(t, f.mod)

	if target, err := os.Readlink(f.dst()); err != nil || target != f.src() {
		t.Errorf("link = %q, %v", target, err)
	}
	if got, err := os.ReadFile(f.dst() + ".eac-bak"); err != nil || string(got) != "locked copy" {
		t.Errorf("backup = %q, %v", got, err)
	}
}
