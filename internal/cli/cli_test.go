package cli

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Automaat/environment-as-code/internal/runner/runnertest"
)

type result struct {
	code           int
	stdout, stderr string
}

func setup(t *testing.T, cfg string) (Env, string) {
	t.Helper()
	root, home := t.TempDir(), t.TempDir()
	files := map[string]string{
		"eac.yaml":       cfg,
		"dotfiles/zshrc": "export EDITOR=vim\n",
	}
	for name, body := range files {
		p := filepath.Join(root, name)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return Env{Home: home, Cwd: root, Runner: runnertest.New()}, home
}

func invoke(env Env, stdin string, args ...string) result {
	var stdout, stderr bytes.Buffer
	env.Args, env.Stdin, env.Stdout, env.Stderr = args, strings.NewReader(stdin), &stdout, &stderr
	code := Main(context.Background(), env)
	return result{code, stdout.String(), stderr.String()}
}

const linksOnly = "links: [{src: dotfiles/zshrc, dst: ~/.zshrc}]\n"

func TestExitCodes(t *testing.T) {
	tests := []struct {
		name string
		cfg  string
		args []string
		want int
		msg  string
	}{
		{"no command", linksOnly, nil, ExitUsage, "Usage"},
		{"unknown command", linksOnly, []string{"destroy"}, ExitUsage, `unknown command "destroy"`},
		{"extra args", linksOnly, []string{"plan", "extra"}, ExitUsage, "Usage"},
		{"unknown module", linksOnly, []string{"--only", "links,nope", "plan"}, ExitUsage, `module "nope"`},
		{"unconfigured module", linksOnly, []string{"--only", "brew", "plan"}, ExitUsage, `module "brew"`},
		{"invalid config", "links: [{src: x}]", []string{"plan"}, ExitErr, "src and dst are required"},
		{"plan with drift", linksOnly, []string{"plan"}, ExitOK, "+ ~/.zshrc"},
		{"check with drift", linksOnly, []string{"check"}, ExitDrift, "+ ~/.zshrc"},
		{"flags after command", linksOnly, []string{"check", "--only", "links"}, ExitDrift, "links:"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env, _ := setup(t, tt.cfg)
			r := invoke(env, "", tt.args...)
			if r.code != tt.want {
				t.Errorf("exit %d, want %d\nstdout: %s\nstderr: %s", r.code, tt.want, r.stdout, r.stderr)
			}
			if !strings.Contains(r.stdout+r.stderr, tt.msg) {
				t.Errorf("output does not mention %q\nstdout: %s\nstderr: %s", tt.msg, r.stdout, r.stderr)
			}
		})
	}
}

func TestApplyConfirmation(t *testing.T) {
	tests := []struct {
		stdin    string
		args     []string
		want     int
		applied  bool
		contains string
	}{
		{"n\n", []string{"apply"}, ExitErr, false, "aborted"},
		{"", []string{"apply"}, ExitErr, false, "aborted"},
		{"yes\n", []string{"apply"}, ExitOK, true, "done"},
		{"", []string{"apply", "-y"}, ExitOK, true, "done"},
	}
	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " ")+" "+strings.TrimSpace(tt.stdin), func(t *testing.T) {
			env, home := setup(t, linksOnly)
			r := invoke(env, tt.stdin, tt.args...)
			if r.code != tt.want || !strings.Contains(r.stdout, tt.contains) {
				t.Fatalf("exit %d, stdout: %s stderr: %s", r.code, r.stdout, r.stderr)
			}
			_, err := os.Lstat(filepath.Join(home, ".zshrc"))
			if applied := err == nil; applied != tt.applied {
				t.Errorf("applied = %v, want %v", applied, tt.applied)
			}
		})
	}
}

func TestApplyNothingToDo(t *testing.T) {
	env, _ := setup(t, linksOnly)
	if r := invoke(env, "", "apply", "-y"); r.code != ExitOK {
		t.Fatal(r)
	}
	r := invoke(env, "", "apply")
	if r.code != ExitOK || !strings.Contains(r.stdout, "nothing to do") {
		t.Errorf("second apply: %+v", r)
	}
	if r := invoke(env, "", "check"); r.code != ExitOK {
		t.Errorf("check after apply: %+v", r)
	}
}

func TestConfigFromEnv(t *testing.T) {
	env, _ := setup(t, linksOnly)
	cfg := filepath.Join(env.Cwd, "eac.yaml")
	env.Cwd = t.TempDir()
	if r := invoke(env, "", "plan"); r.code != ExitErr {
		t.Fatalf("without config: %+v", r)
	}
	t.Setenv("EAC_CONFIG", cfg)
	if r := invoke(env, "", "plan"); r.code != ExitOK {
		t.Errorf("with EAC_CONFIG: %+v", r)
	}
	if r := invoke(env, "", "-c", cfg, "plan"); r.code != ExitOK {
		t.Errorf("with -c: %+v", r)
	}
}

type brokenWriter struct{}

func (brokenWriter) Write([]byte) (int, error) { return 0, os.ErrClosed }

func TestOutputFailureFailsTheRun(t *testing.T) {
	env, _ := setup(t, linksOnly)
	env.Args, env.Stdin, env.Stdout, env.Stderr = []string{"plan"}, strings.NewReader(""), brokenWriter{}, &bytes.Buffer{}
	if code := Main(context.Background(), env); code != ExitErr {
		t.Errorf("exit %d, want %d when stdout is broken", code, ExitErr)
	}
}
