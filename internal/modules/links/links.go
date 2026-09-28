// Package links symlinks repo files into the home directory.
package links

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
)

type Module struct {
	Links []config.Link
	Paths config.Paths
	Now   func() time.Time
}

func (m *Module) Name() string { return "links" }

func (m *Module) Plan(_ context.Context) ([]engine.Change, error) {
	var changes []engine.Change
	for _, l := range m.Links {
		src, dst := m.Paths.Src(l.Src), m.Paths.Dst(l.Dst)
		c, err := m.plan(src, dst)
		if err != nil {
			return nil, err
		}
		if c != nil {
			changes = append(changes, *c)
		}
	}
	return changes, nil
}

func (m *Module) plan(src, dst string) (*engine.Change, error) {
	if _, err := os.Stat(src); err != nil {
		return nil, fmt.Errorf("source %s: %w", src, err)
	}
	target := m.Paths.Pretty(dst)
	link := func(ctx context.Context) error { return symlink(src, dst) }

	info, err := os.Lstat(dst)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return &engine.Change{Action: engine.Create, Target: target, Detail: "→ " + m.Paths.Pretty(src), Apply: link}, nil
	case err != nil:
		return nil, err
	}

	if info.Mode()&fs.ModeSymlink != 0 {
		current, err := os.Readlink(dst)
		if err != nil {
			return nil, err
		}
		if current == src {
			return nil, nil
		}
		return &engine.Change{
			Action: engine.Update, Target: target,
			Detail: fmt.Sprintf("relink %s → %s", current, m.Paths.Pretty(src)),
			Apply: func(ctx context.Context) error {
				if err := os.Remove(dst); err != nil {
					return err
				}
				return symlink(src, dst)
			},
		}, nil
	}

	if info.IsDir() {
		return nil, fmt.Errorf("%s is a directory; move it away before linking", dst)
	}
	backup := m.backupPath(dst)
	return &engine.Change{
		Action: engine.Update, Target: target,
		Detail: fmt.Sprintf("back up to %s, link → %s", filepath.Base(backup), m.Paths.Pretty(src)),
		Apply: func(ctx context.Context) error {
			if err := os.Rename(dst, backup); err != nil {
				return err
			}
			return symlink(src, dst)
		},
	}, nil
}

func (m *Module) backupPath(dst string) string {
	p := dst + ".eac-bak"
	if _, err := os.Lstat(p); errors.Is(err, fs.ErrNotExist) {
		return p
	}
	return fmt.Sprintf("%s.%d", p, m.now().Unix())
}

func (m *Module) now() time.Time {
	if m.Now != nil {
		return m.Now()
	}
	return time.Now()
}

func symlink(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	return os.Symlink(src, dst)
}
