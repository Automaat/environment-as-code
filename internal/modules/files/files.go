// Package files installs repo files into the home directory as protected,
// read-only copies. Changes go through the repo and `eac apply`.
package files

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
	"github.com/Automaat/environment-as-code/internal/install"
)

// Module installs Files. Keep lists destinations the shared installer writes
// for other modules (templates), so they aren't treated as orphans.
type Module struct {
	Files     []config.Link
	Keep      []string
	Paths     config.Paths
	Installer *install.Installer
}

func (m *Module) Name() string { return "files" }

func (m *Module) Plan(_ context.Context) ([]engine.Change, error) {
	wanted := map[string]bool{}
	for _, k := range m.Keep {
		wanted[k] = true
	}
	var changes []engine.Change
	for _, f := range m.Files {
		pairs, err := expand(m.Paths.Src(f.Src), m.Paths.Dst(f.Dst))
		if err != nil {
			return nil, err
		}
		for _, p := range pairs {
			wanted[p.dst] = true
			c, err := m.plan(p)
			if err != nil {
				return nil, err
			}
			if c != nil {
				changes = append(changes, *c)
			}
		}
	}
	for _, dst := range m.Installer.State.Keys() {
		if wanted[dst] {
			continue
		}
		c, err := m.Installer.PlanRemove(dst)
		if err != nil {
			return nil, err
		}
		changes = append(changes, *c)
	}
	return changes, nil
}

func (m *Module) plan(p pair) (*engine.Change, error) {
	info, err := os.Stat(p.src)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(p.src)
	if err != nil {
		return nil, err
	}
	return m.Installer.Plan(p.dst, data, info.Mode().Perm())
}

type pair struct{ src, dst string }

// expand maps a source file to itself, or a source directory to every file
// inside it. Extra files already in a destination directory are left alone.
func expand(src, dst string) ([]pair, error) {
	info, err := os.Stat(src)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return []pair{{src, dst}}, nil
	}
	var pairs []pair
	err = filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		pairs = append(pairs, pair{path, filepath.Join(dst, rel)})
		return nil
	})
	return pairs, err
}
