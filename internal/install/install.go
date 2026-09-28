// Package install writes protected copies of managed files: write bits
// stripped and, optionally, the macOS immutable flag set, so dotfiles can only
// change through the repo and `eac apply`.
package install

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	udiff "github.com/aymanbagabas/go-udiff"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
)

type Installer struct {
	Paths     config.Paths
	State     *State
	Immutable bool
	Now       func() time.Time
}

// ReadOnly strips write bits, keeping read and execute bits of perm.
func ReadOnly(perm fs.FileMode) fs.FileMode {
	return perm.Perm() &^ 0o222
}

// Plan returns the change that makes dst a protected copy of want, or nil.
func (in *Installer) Plan(dst string, want []byte, perm fs.FileMode) (*engine.Change, error) {
	perm = ReadOnly(perm)
	target := in.Paths.Pretty(dst)
	change := func(a engine.Action, detail string, apply func(context.Context) error) *engine.Change {
		return &engine.Change{Action: a, Target: target, Detail: detail, Apply: apply}
	}

	info, err := os.Lstat(dst)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		c := change(engine.Create, "", in.install(dst, want, perm, ""))
		c.Diff = udiff.Unified("/dev/null", "repo", "", string(want))
		return c, nil
	case err != nil:
		return nil, err
	case info.IsDir():
		return nil, fmt.Errorf("%s is a directory; move it away first", dst)
	case info.Mode()&fs.ModeSymlink != 0:
		c := change(engine.Update, "replace symlink", in.install(dst, want, perm, ""))
		linked, _ := os.ReadFile(dst)
		c.Diff = udiff.Unified(dst, "repo", string(linked), string(want))
		return c, nil
	}

	have, err := os.ReadFile(dst)
	if err != nil {
		return nil, err
	}
	withDiff := func(c *engine.Change) *engine.Change {
		c.Diff = udiff.Unified(dst, "repo", string(have), string(want))
		return c
	}
	if bytes.Equal(have, want) {
		reasons, err := in.protectionDrift(dst, info.Mode().Perm(), perm)
		if err != nil || len(reasons) == 0 {
			return nil, err
		}
		return change(engine.Update, strings.Join(reasons, ", "), in.protect(dst, want, perm)), nil
	}

	recorded, known := in.State.Get(dst)
	if known && recorded == Sum(have) {
		return withDiff(change(engine.Update, "content", in.install(dst, want, perm, ""))), nil
	}
	backup := BackupPath(dst, in.now())
	why := "edited in place"
	if !known {
		why = "not managed yet"
	}
	return withDiff(change(engine.Update, fmt.Sprintf("%s, back up to %s", why, filepath.Base(backup)), in.install(dst, want, perm, backup))), nil
}

// PlanRemove handles a destination eac once wrote but no longer manages. A
// copy it still recognizes is deleted, an edited one is backed up first, and
// anything that is no longer a regular file (e.g. now a link another module
// owns) is only dropped from the state.
func (in *Installer) PlanRemove(dst string) (*engine.Change, error) {
	target := in.Paths.Pretty(dst)
	forget := func(context.Context) error { return in.State.Forget(dst) }

	info, err := os.Lstat(dst)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return &engine.Change{Action: engine.Remove, Target: target, Detail: "already gone, forget", Apply: forget}, nil
	case err != nil:
		return nil, err
	case !info.Mode().IsRegular():
		return &engine.Change{Action: engine.Remove, Target: target, Detail: "no longer managed, left in place", Apply: forget}, nil
	}

	have, err := os.ReadFile(dst)
	if err != nil {
		return nil, err
	}
	if recorded, _ := in.State.Get(dst); recorded == Sum(have) {
		return &engine.Change{
			Action: engine.Remove, Target: target, Detail: "no longer managed",
			Apply: func(ctx context.Context) error {
				if err := unlock(dst); err != nil {
					return err
				}
				if err := os.Remove(dst); err != nil {
					return err
				}
				return forget(ctx)
			},
		}, nil
	}
	backup := BackupPath(dst, in.now())
	return &engine.Change{
		Action: engine.Remove, Target: target,
		Detail: "no longer managed, edited: back up to " + filepath.Base(backup),
		Apply: func(ctx context.Context) error {
			if err := unlock(dst); err != nil {
				return err
			}
			if err := os.Rename(dst, backup); err != nil {
				return err
			}
			if err := os.Chmod(backup, 0o644); err != nil {
				return err
			}
			return forget(ctx)
		},
	}, nil
}

func (in *Installer) protectionDrift(dst string, have, want fs.FileMode) ([]string, error) {
	var reasons []string
	if have != want {
		reasons = append(reasons, fmt.Sprintf("mode %o → %o", have, want))
	}
	if !immutableSupported {
		return reasons, nil
	}
	locked, err := isImmutable(dst)
	if err != nil {
		return nil, err
	}
	switch {
	case in.Immutable && !locked:
		reasons = append(reasons, "set immutable")
	case !in.Immutable && locked:
		reasons = append(reasons, "clear immutable")
	}
	return reasons, nil
}

// install writes want to dst, first moving any existing file to backup when
// one is given. The immutable flag is lifted only for the swap.
func (in *Installer) install(dst string, want []byte, perm fs.FileMode, backup string) func(context.Context) error {
	return func(context.Context) error {
		if err := unlock(dst); err != nil {
			return err
		}
		if backup != "" {
			if err := os.Rename(dst, backup); err != nil {
				return err
			}
			if err := os.Chmod(backup, 0o644); err != nil {
				return err
			}
		}
		if err := writeAtomic(dst, want, perm); err != nil {
			return err
		}
		return in.finish(dst, want)
	}
}

func (in *Installer) protect(dst string, want []byte, perm fs.FileMode) func(context.Context) error {
	return func(context.Context) error {
		if err := unlock(dst); err != nil {
			return err
		}
		if err := os.Chmod(dst, perm); err != nil {
			return err
		}
		return in.finish(dst, want)
	}
}

func (in *Installer) finish(dst string, want []byte) error {
	if in.Immutable {
		if err := setImmutable(dst, true); err != nil {
			return err
		}
	}
	return in.State.Record(dst, Sum(want))
}

func (in *Installer) now() time.Time {
	if in.Now != nil {
		return in.Now()
	}
	return time.Now()
}

// unlock clears the immutable flag on an existing regular file so it can be
// replaced; symlinks and missing files need nothing.
func unlock(path string) error {
	info, err := os.Lstat(path)
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	if !info.Mode().IsRegular() {
		return nil
	}
	return setImmutable(path, false)
}

// Unlock clears the immutable flag on every file under root, for tests and for
// tearing down a managed tree.
func Unlock(root string) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		return unlock(path)
	})
}

// BackupPath picks a free backup name next to dst.
func BackupPath(dst string, now time.Time) string {
	p := dst + ".eac-bak"
	if _, err := os.Lstat(p); errors.Is(err, fs.ErrNotExist) {
		return p
	}
	return fmt.Sprintf("%s.%d", p, now.Unix())
}

// writeAtomic replaces dst via rename so a failed write never leaves a
// half-written file behind.
func writeAtomic(dst string, data []byte, perm fs.FileMode) (err error) {
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
	if err = os.Chmod(tmp.Name(), perm); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), dst)
}
