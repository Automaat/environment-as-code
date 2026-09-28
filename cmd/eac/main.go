package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Automaat/environment-as-code/internal/cli"
	"github.com/Automaat/environment-as-code/internal/runner"
)

func main() {
	os.Exit(run())
}

func run() int {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

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
	return cli.Main(ctx, cli.Env{
		Args:    os.Args[1:],
		Stdin:   os.Stdin,
		Stdout:  os.Stdout,
		Stderr:  os.Stderr,
		Home:    home,
		Cwd:     cwd,
		Runner:  runner.NewExec(),
		PAMFile: os.Getenv("EAC_PAM_FILE"),
	})
}
