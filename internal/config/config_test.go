package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeConfig(t *testing.T, body string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), FileName)
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestLoad(t *testing.T) {
	p := writeConfig(t, `
links:
  - {src: dotfiles/zsh/.zshrc, dst: ~/.zshrc}
templates:
  vars: {theme: nord}
  files:
    - {src: a.tmpl, dst: ~/a, mode: 0600}
brew: {file: Brewfile, cleanup: zap, upgrade: true}
mise: {config: dotfiles/mise/config.toml}
defaults:
  - {domain: com.apple.dock, key: autohide, value: true}
  - {domain: NSGlobalDomain, key: KeyRepeat, value: 2}
  - {domain: x, key: f, value: 1.5}
  - {domain: x, key: s, value: Dark}
system:
  dirs: [{path: ~/.ssh, mode: 0700}]
  sudoTouchID: true
  sshKey: {path: ~/.ssh/id_ed25519, comment: me}
commands:
  - {name: n, check: "true", run: "true"}
`)
	c, err := Load(p)
	if err != nil {
		t.Fatal(err)
	}
	if want, _ := filepath.EvalSymlinks(filepath.Dir(p)); c.Root != want {
		t.Errorf("Root = %q, want %q", c.Root, want)
	}
	if got := c.Templates.Files[0].Mode; got != 0o600 {
		t.Errorf("template mode = %o, want 600", got)
	}
	if got := c.System.Dirs[0].Mode; got != 0o700 {
		t.Errorf("dir mode = %o, want 700", got)
	}
	wantTypes := []any{true, 2, 1.5, "Dark"}
	for i, want := range wantTypes {
		if c.Defaults[i].Value != want {
			t.Errorf("defaults[%d].Value = %#v, want %#v", i, c.Defaults[i].Value, want)
		}
	}
}

func TestLoadResolvesSymlinkedRoot(t *testing.T) {
	real := filepath.Dir(writeConfig(t, linksOnly))
	alias := filepath.Join(t.TempDir(), "alias")
	if err := os.Symlink(real, alias); err != nil {
		t.Fatal(err)
	}
	c, err := Load(filepath.Join(alias, FileName))
	if err != nil {
		t.Fatal(err)
	}
	want, _ := filepath.EvalSymlinks(real)
	if c.Root != want {
		t.Errorf("Root = %q, want %q", c.Root, want)
	}
}

const linksOnly = "links: [{src: a, dst: ~/a}]\n"

func TestLoadRejects(t *testing.T) {
	tests := []struct {
		name string
		body string
		want []string
	}{
		{"unknown field", "linkz: []", []string{"field linkz not found"}},
		{"missing link fields", "links: [{src: a}]", []string{"links[0]: src and dst are required"}},
		{"duplicate destination", `
links: [{src: a, dst: ~/x}]
templates: {files: [{src: b, dst: ~/x}]}`, []string{`destination "~/x" is managed twice`}},
		{"bad cleanup", "brew: {file: B, cleanup: nuke}", []string{`brew.cleanup: "nuke"`}},
		{"brew without file", "brew: {upgrade: true}", []string{"brew.file is required"}},
		{"unsupported default", "defaults: [{domain: d, key: k, value: [1]}]", []string{"unsupported value"}},
		{"reports all errors", `
links: [{src: a}]
commands: [{name: n}]`, []string{"links[0]", "commands[0]"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Load(writeConfig(t, tt.body))
			if err == nil {
				t.Fatal("expected error")
			}
			for _, w := range tt.want {
				if !strings.Contains(err.Error(), w) {
					t.Errorf("error %q does not mention %q", err, w)
				}
			}
		})
	}
}

func TestFind(t *testing.T) {
	root := t.TempDir()
	deep := filepath.Join(root, "a", "b")
	if err := os.MkdirAll(deep, 0o755); err != nil {
		t.Fatal(err)
	}
	if _, err := Find(deep); err == nil {
		t.Fatal("expected not found")
	}
	want := filepath.Join(root, FileName)
	if err := os.WriteFile(want, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := Find(deep)
	if err != nil || got != want {
		t.Fatalf("Find = %q, %v; want %q", got, err, want)
	}
}

func TestPaths(t *testing.T) {
	p := Paths{Home: "/h", Root: "/r"}
	tests := []struct {
		fn   func(string) string
		in   string
		want string
	}{
		{p.Dst, "~", "/h"},
		{p.Dst, "~/.zshrc", "/h/.zshrc"},
		{p.Dst, "/etc/x", "/etc/x"},
		{p.Dst, "~user/x", "~user/x"},
		{p.Src, "dotfiles/a", "/r/dotfiles/a"},
		{p.Src, "/abs", "/abs"},
		{p.Pretty, "/h/.zshrc", "~/.zshrc"},
		{p.Pretty, "/hx/y", "/hx/y"},
		{p.Pretty, "/r/a", "/r/a"},
	}
	for _, tt := range tests {
		if got := tt.fn(tt.in); got != tt.want {
			t.Errorf("(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
