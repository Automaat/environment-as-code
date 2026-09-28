package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Automaat/environment-as-code/internal/cli"
	"github.com/Automaat/environment-as-code/internal/runner"
)

// exitInterrupted follows the shell convention of 128 + SIGINT.
const exitInterrupted = 130

func main() {
	os.Exit(run())
}

// run restores default SIGINT handling after the first one, so a second
// Ctrl-C kills eac instead of being swallowed.
func run() int {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		<-ctx.Done()
		stop()
	}()

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "eac:", err)
		return cli.ExitErr
	}
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "eac:", err)
		return cli.ExitErr
	}
	code := cli.Main(ctx, cli.Env{
		Args:    os.Args[1:],
		Stdin:   os.Stdin,
		Stdout:  os.Stdout,
		Stderr:  os.Stderr,
		Home:    home,
		Cwd:     cwd,
		Runner:  runner.NewExec(),
		PAMFile: os.Getenv("EAC_PAM_FILE"),
	})
	if ctx.Err() != nil {
		return exitInterrupted
	}
	return code
}
