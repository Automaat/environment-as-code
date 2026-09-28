// Package templates renders repo templates into real files, for configs that
// need machine values (home path) or must not be symlinks (gpg).
package templates

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
)

const defaultMode fs.FileMode = 0o644

// Data is what templates can reference.
type Data struct {
	Home string
	Vars map[string]string
}

type Module struct {
	Templates config.Templates
	Paths     config.Paths
}

func (m *Module) Name() string { return "templates" }

func (m *Module) Plan(_ context.Context) ([]engine.Change, error) {
	data := Data{Home: m.Paths.Home, Vars: m.Templates.Vars}
	var changes []engine.Change
	for _, f := range m.Templates.Files {
		src, dst := m.Paths.Src(f.Src), m.Paths.Dst(f.Dst)
		mode := f.Mode
		if mode == 0 {
			mode = defaultMode
		}
		want, err := Render(src, data)
		if err != nil {
			return nil, err
		}
		c, err := m.plan(dst, want, mode)
		if err != nil {
			return nil, err
		}
		if c != nil {
			changes = append(changes, *c)
		}
	}
	return changes, nil
}

// Render executes the template at path, failing on unknown variables.
func Render(path string, data Data) ([]byte, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	t, err := template.New(filepath.Base(path)).Option("missingkey=error").Parse(string(raw))
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (m *Module) plan(dst string, want []byte, mode fs.FileMode) (*engine.Change, error) {
	target := m.Paths.Pretty(dst)
	write := func(ctx context.Context) error { return writeFile(dst, want, mode) }

	info, err := os.Lstat(dst)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return &engine.Change{Action: engine.Create, Target: target, Apply: write}, nil
	case err != nil:
		return nil, err
	case info.IsDir():
		return nil, fmt.Errorf("%s is a directory", dst)
	case info.Mode()&fs.ModeSymlink != 0:
		return &engine.Change{Action: engine.Update, Target: target, Detail: "replace symlink with rendered file", Apply: write}, nil
	}

	have, err := os.ReadFile(dst)
	if err != nil {
		return nil, err
	}
	var reasons []string
	if !bytes.Equal(have, want) {
		reasons = append(reasons, "content")
	}
	if info.Mode().Perm() != mode {
		reasons = append(reasons, fmt.Sprintf("mode %o → %o", info.Mode().Perm(), mode))
	}
	if len(reasons) == 0 {
		return nil, nil
	}
	return &engine.Change{Action: engine.Update, Target: target, Detail: strings.Join(reasons, ", "), Apply: write}, nil
}

// writeFile replaces dst atomically so a failed write never leaves a
// half-written config behind.
func writeFile(dst string, data []byte, mode fs.FileMode) (err error) {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(dst), ".eac-*")
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err = errors.Join(err, os.Remove(tmp.Name()))
		}
	}()
	_, err = tmp.Write(data)
	if err = errors.Join(err, tmp.Close()); err != nil {
		return err
	}
	if err = os.Chmod(tmp.Name(), mode); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), dst)
}
