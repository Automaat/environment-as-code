package files

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
	"github.com/Automaat/environment-as-code/internal/install"
)

func TestFilesAndDirectories(t *testing.T) {
	home, root := t.TempDir(), t.TempDir()
	write := func(rel, body string, mode os.FileMode) {
		p := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), mode); err != nil {
			t.Fatal(err)
		}
	}
	write("zsh/zshrc", "zsh", 0o644)
	write("bin/tool", "#!/bin/sh", 0o755)
	write("ghostty/themes/nord", "a", 0o644)
	write("ghostty/themes/sub/dark", "b", 0o644)

	state, err := install.LoadState(install.StatePath(home))
	if err != nil {
		t.Fatal(err)
	}
	paths := config.Paths{Home: home, Root: root}
	m := &Module{
		Paths:     paths,
		Installer: &install.Installer{Paths: paths, State: state},
		Files: []config.Link{
			{Src: "zsh/zshrc", Dst: "~/.zshrc"},
			{Src: "bin/tool", Dst: "~/.local/bin/tool"},
			{Src: "ghostty/themes", Dst: "~/.config/ghostty/themes"},
		},
	}
	ctx := context.Background()

	changes, err := m.Plan(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(changes) != 4 {
		t.Fatalf("changes = %v", changes)
	}
	if err := engine.Apply(ctx, func(string) {}, engine.Plan{{Changes: changes}}); err != nil {
		t.Fatal(err)
	}

	want := map[string]os.FileMode{
		".zshrc":                          0o444,
		".local/bin/tool":                 0o555,
		".config/ghostty/themes/nord":     0o444,
		".config/ghostty/themes/sub/dark": 0o444,
	}
	for rel, mode := range want {
		info, err := os.Lstat(filepath.Join(home, rel))
		if err != nil || !info.Mode().IsRegular() || info.Mode().Perm() != mode {
			t.Errorf("%s: %v, %v; want regular file %o", rel, info, err, mode)
		}
	}
	if again, err := m.Plan(ctx); err != nil || len(again) != 0 {
		t.Errorf("not idempotent: %v, %v", again, err)
	}
}

func TestMissingSource(t *testing.T) {
	home := t.TempDir()
	state, _ := install.LoadState(install.StatePath(home))
	paths := config.Paths{Home: home, Root: t.TempDir()}
	m := &Module{
		Paths: paths, Installer: &install.Installer{Paths: paths, State: state},
		Files: []config.Link{{Src: "nope", Dst: "~/x"}},
	}
	if _, err := m.Plan(context.Background()); err == nil {
		t.Error("expected error")
	}
}

func TestRemovesFilesNoLongerManaged(t *testing.T) {
	home, root := t.TempDir(), t.TempDir()
	t.Cleanup(func() { _ = install.Unlock(home) })
	for rel, body := range map[string]string{"zshrc": "z", "vimrc": "v", "themes/a": "a", "themes/b": "b"} {
		p := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	state, err := install.LoadState(install.StatePath(home))
	if err != nil {
		t.Fatal(err)
	}
	paths := config.Paths{Home: home, Root: root}
	in := &install.Installer{Paths: paths, State: state, Immutable: true}
	template := filepath.Join(home, ".config/k9s/config.yaml")
	if c, err := in.Plan(template, []byte("rendered"), 0o644); err != nil || c.Apply(context.Background()) != nil {
		t.Fatalf("template setup: %v", err)
	}
	m := &Module{Paths: paths, Installer: in, Keep: []string{template}, Files: []config.Link{
		{Src: "zshrc", Dst: "~/.zshrc"},
		{Src: "vimrc", Dst: "~/.vimrc"},
		{Src: "themes", Dst: "~/.themes"},
	}}
	apply := func() []engine.Change {
		changes, err := m.Plan(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if err := engine.Apply(context.Background(), func(string) {}, engine.Plan{{Changes: changes}}); err != nil {
			t.Fatal(err)
		}
		return changes
	}
	apply()

	m.Files = m.Files[:1]
	m.Files = append(m.Files, config.Link{Src: "themes", Dst: "~/.themes"})
	if err := os.Remove(filepath.Join(root, "themes/b")); err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, c := range apply() {
		got = append(got, string(c.Action)+" "+c.Target)
	}
	want := []string{"- ~/.themes/b", "- ~/.vimrc"}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("changes = %v, want %v", got, want)
	}
	for _, rel := range []string{".vimrc", ".themes/b"} {
		if _, err := os.Lstat(filepath.Join(home, rel)); err == nil {
			t.Errorf("%s still exists", rel)
		}
	}
	if _, err := os.Stat(template); err != nil {
		t.Errorf("template output must survive: %v", err)
	}
	if again, _ := m.Plan(context.Background()); len(again) != 0 {
		t.Errorf("not converged: %v", again)
	}
}
