package install

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
)

func newInstaller(t *testing.T, immutable bool) (*Installer, string) {
	t.Helper()
	home := t.TempDir()
	t.Cleanup(func() {
		if err := Unlock(home); err != nil {
			t.Errorf("unlock: %v", err)
		}
	})
	state, err := LoadState(StatePath(home))
	if err != nil {
		t.Fatal(err)
	}
	return &Installer{
		Paths:     config.Paths{Home: home},
		State:     state,
		Immutable: immutable,
		Now:       func() time.Time { return time.Unix(1700000000, 0) },
	}, home
}

func plan(t *testing.T, in *Installer, dst, want string, perm os.FileMode) *engine.Change {
	t.Helper()
	c, err := in.Plan(dst, []byte(want), perm)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func converge(t *testing.T, in *Installer, dst, want string, perm os.FileMode) *engine.Change {
	t.Helper()
	c := plan(t, in, dst, want, perm)
	if c == nil {
		return nil
	}
	if err := c.Apply(context.Background()); err != nil {
		t.Fatal(err)
	}
	if again := plan(t, in, dst, want, perm); again != nil {
		t.Fatalf("not idempotent: %s", again)
	}
	return c
}

func mustWrite(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func assertProtected(t *testing.T, path, body string, perm os.FileMode, immutable bool) {
	t.Helper()
	got, err := os.ReadFile(path)
	if err != nil || string(got) != body {
		t.Errorf("content = %q, %v; want %q", got, err, body)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != perm {
		t.Errorf("mode = %o, want %o", info.Mode().Perm(), perm)
	}
	if locked, _ := isImmutable(path); immutableSupported && locked != immutable {
		t.Errorf("immutable = %v, want %v", locked, immutable)
	}
}

func TestReadOnly(t *testing.T) {
	for in, want := range map[os.FileMode]os.FileMode{0o644: 0o444, 0o755: 0o555, 0o600: 0o400, 0o444: 0o444} {
		if got := ReadOnly(in); got != want {
			t.Errorf("ReadOnly(%o) = %o, want %o", in, got, want)
		}
	}
}

func TestCreate(t *testing.T) {
	for _, immutable := range []bool{false, true} {
		in, home := newInstaller(t, immutable)
		dst := filepath.Join(home, ".config/git/config")

		c := converge(t, in, dst, "v1", 0o644)

		if c.Action != engine.Create || c.Target != "~/.config/git/config" {
			t.Errorf("change = %s", c)
		}
		assertProtected(t, dst, "v1", 0o444, immutable)
		if sum, _ := in.State.Get(dst); sum != Sum([]byte("v1")) {
			t.Error("state not recorded")
		}
		reloaded, err := LoadState(StatePath(home))
		if err != nil || reloaded.Files[dst] != Sum([]byte("v1")) {
			t.Errorf("state not persisted: %v %v", reloaded, err)
		}
	}
}

func TestWriteIsBlocked(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("root ignores permissions")
	}
	in, home := newInstaller(t, true)
	dst := filepath.Join(home, ".zshrc")
	converge(t, in, dst, "v1", 0o644)

	if err := os.WriteFile(dst, []byte("edited"), 0o644); err == nil {
		t.Error("write to a protected file succeeded")
	}
	if immutableSupported {
		if err := os.Chmod(dst, 0o644); err == nil {
			t.Error("chmod of an immutable file succeeded")
		}
		if err := os.Remove(dst); err == nil {
			t.Error("delete of an immutable file succeeded")
		}
	}
}

// edit simulates a user forcing an in-place change past the protection.
func edit(t *testing.T, path, body string) {
	t.Helper()
	if err := unlock(path); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(path, 0o644); err != nil {
		t.Fatal(err)
	}
	mustWrite(t, path, body)
}

func TestUpdates(t *testing.T) {
	tests := []struct {
		name       string
		setup      func(t *testing.T, in *Installer, dst string)
		wantDetail string
		wantBackup string
	}{
		{
			name:       "repo changed: overwrite without backup",
			setup:      func(t *testing.T, in *Installer, dst string) { converge(t, in, dst, "old", 0o644) },
			wantDetail: "content",
		},
		{
			name: "edited in place: back up the edit",
			setup: func(t *testing.T, in *Installer, dst string) {
				converge(t, in, dst, "old", 0o644)
				edit(t, dst, "my edit")
			},
			wantDetail: "edited in place, back up to .zshrc.eac-bak",
			wantBackup: "my edit",
		},
		{
			name:       "pre-existing unmanaged file: back it up",
			setup:      func(t *testing.T, _ *Installer, dst string) { mustWrite(t, dst, "hand written") },
			wantDetail: "not managed yet, back up to .zshrc.eac-bak",
			wantBackup: "hand written",
		},
		{
			name: "home-manager symlink: replace",
			setup: func(t *testing.T, _ *Installer, dst string) {
				if err := os.Symlink("/nix/store/x-home-manager-files/.zshrc", dst); err != nil {
					t.Fatal(err)
				}
			},
			wantDetail: "replace symlink",
		},
		{
			name: "protection removed: restore it",
			setup: func(t *testing.T, in *Installer, dst string) {
				converge(t, in, dst, "new", 0o644)
				if err := unlock(dst); err != nil {
					t.Fatal(err)
				}
				if err := os.Chmod(dst, 0o644); err != nil {
					t.Fatal(err)
				}
			},
			wantDetail: "mode 644 → 444",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, home := newInstaller(t, true)
			dst := filepath.Join(home, ".zshrc")
			tt.setup(t, in, dst)

			c := converge(t, in, dst, "new", 0o644)

			if c == nil || c.Action != engine.Update || !strings.HasPrefix(c.Detail, tt.wantDetail) {
				t.Fatalf("change = %v, want detail %q", c, tt.wantDetail)
			}
			assertProtected(t, dst, "new", 0o444, true)
			backup, err := os.ReadFile(dst + ".eac-bak")
			switch {
			case tt.wantBackup == "" && err == nil:
				t.Errorf("unexpected backup %q", backup)
			case tt.wantBackup != "" && string(backup) != tt.wantBackup:
				t.Errorf("backup = %q, %v; want %q", backup, err, tt.wantBackup)
			}
		})
	}
}

func TestImmutableToggle(t *testing.T) {
	if !immutableSupported {
		t.Skip("no immutable flag on this OS")
	}
	in, home := newInstaller(t, true)
	dst := filepath.Join(home, "f")
	converge(t, in, dst, "x", 0o644)

	in.Immutable = false
	c := converge(t, in, dst, "x", 0o644)
	if c == nil || c.Detail != "clear immutable" {
		t.Fatalf("change = %v", c)
	}
	assertProtected(t, dst, "x", 0o444, false)
}

func TestExecutableBitKept(t *testing.T) {
	in, home := newInstaller(t, false)
	dst := filepath.Join(home, ".local/bin/tool")
	converge(t, in, dst, "#!/bin/sh\n", 0o755)
	assertProtected(t, dst, "#!/bin/sh\n", 0o555, false)
}

func TestDirectoryInTheWay(t *testing.T) {
	in, home := newInstaller(t, false)
	dst := filepath.Join(home, "d")
	if err := os.Mkdir(dst, 0o755); err != nil {
		t.Fatal(err)
	}
	if _, err := in.Plan(dst, nil, 0o644); err == nil {
		t.Error("expected error")
	}
}

func TestBackupPathAvoidsCollisions(t *testing.T) {
	dir := t.TempDir()
	dst := filepath.Join(dir, "f")
	now := time.Unix(42, 0)
	if got := BackupPath(dst, now); got != dst+".eac-bak" {
		t.Errorf("first = %s", got)
	}
	mustWrite(t, dst+".eac-bak", "")
	if got := BackupPath(dst, now); got != dst+".eac-bak.42" {
		t.Errorf("second = %s", got)
	}
}

func TestLoadStateCorrupt(t *testing.T) {
	p := filepath.Join(t.TempDir(), "files.json")
	mustWrite(t, p, "{nope")
	if _, err := LoadState(p); err == nil {
		t.Error("expected error for corrupt state")
	}
}
