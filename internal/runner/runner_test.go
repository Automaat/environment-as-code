package runner

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestExecRun(t *testing.T) {
	ctx := context.Background()
	e := NewExec()

	res, err := e.Run(ctx, Cmd{Name: "sh", Args: []string{"-c", `printf "$FOO"; cat; echo err >&2; exit 3`}, Env: []string{"FOO=hi-"}, Stdin: "in"})
	if err != nil {
		t.Fatal(err)
	}
	if res.Stdout != "hi-in" || strings.TrimSpace(res.Stderr) != "err" || res.ExitCode != 3 {
		t.Errorf("res = %+v", res)
	}

	if _, err := e.Run(ctx, Cmd{Name: "definitely-not-a-binary-eac"}); err == nil {
		t.Error("expected start error")
	}
}

func TestExecRunDir(t *testing.T) {
	dir := t.TempDir()
	out, err := Output(context.Background(), NewExec(), Cmd{Name: "pwd", Dir: dir})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasSuffix(strings.TrimSpace(out), dir) {
		t.Errorf("pwd = %q, want suffix %q", out, dir)
	}
}

func TestExecStream(t *testing.T) {
	var stdout, stderr bytes.Buffer
	e := &Exec{Stdout: &stdout, Stderr: &stderr}
	res, err := e.Run(context.Background(), Cmd{Name: "sh", Args: []string{"-c", "echo out; echo err >&2"}, Stream: true})
	if err != nil {
		t.Fatal(err)
	}
	if res.Stdout != "" || stdout.String() != "out\n" || stderr.String() != "err\n" {
		t.Errorf("res = %+v, stdout %q, stderr %q", res, stdout.String(), stderr.String())
	}
}

func TestResultErr(t *testing.T) {
	c := Cmd{Name: "x", Args: []string{"y"}}
	if err := (Result{}).Err(c); err != nil {
		t.Errorf("zero exit: %v", err)
	}
	tests := []struct {
		res  Result
		want string
	}{
		{Result{ExitCode: 2, Stderr: "bad\n"}, "x y: exit 2: bad"},
		{Result{ExitCode: 1, Stdout: "only stdout"}, "x y: exit 1: only stdout"},
	}
	for _, tt := range tests {
		if err := tt.res.Err(c); err == nil || err.Error() != tt.want {
			t.Errorf("Err = %v, want %q", err, tt.want)
		}
	}
}
