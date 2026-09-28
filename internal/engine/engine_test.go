package engine

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"
)

type stubModule struct {
	name    string
	changes []Change
	err     error
}

func (s stubModule) Name() string                           { return s.name }
func (s stubModule) Plan(context.Context) ([]Change, error) { return s.changes, s.err }

func TestBuildStopsOnPlanError(t *testing.T) {
	_, err := Build(context.Background(), []Module{
		stubModule{name: "ok"},
		stubModule{name: "bad", err: errors.New("boom")},
	})
	if err == nil || !strings.Contains(err.Error(), "plan bad: boom") {
		t.Fatalf("err = %v", err)
	}
}

func TestApply(t *testing.T) {
	var ran []string
	record := func(name string, err error) func(context.Context) error {
		return func(context.Context) error {
			ran = append(ran, name)
			return err
		}
	}
	plan := Plan{
		{Module: "a", Changes: []Change{
			{Action: Create, Target: "info-only"},
			{Action: Run, Target: "a1", Apply: record("a1", nil)},
			{Action: Run, Target: "a2", Apply: record("a2", errors.New("fail"))},
			{Action: Run, Target: "a3", Apply: record("a3", nil)},
		}},
		{Module: "b", Changes: []Change{
			{Action: Run, Target: "b1", Apply: record("b1", nil)},
		}},
	}

	var progress []string
	err := Apply(context.Background(), func(s string) { progress = append(progress, s) }, plan)

	if got, want := strings.Join(ran, ","), "a1,a2,b1"; got != want {
		t.Errorf("ran %s, want %s (module stops at first failure, others continue)", got, want)
	}
	if err == nil || !strings.Contains(err.Error(), "a: a2: fail") {
		t.Errorf("err = %v", err)
	}
	if got, want := strings.Join(progress, "|"), "a: ! a1|a: ! a2|b: ! b1"; got != want {
		t.Errorf("progress %s, want %s", got, want)
	}
}

func TestPrint(t *testing.T) {
	var out bytes.Buffer
	err := Print(&out, Plan{
		{Module: "links", Changes: []Change{{Action: Create, Target: "~/.zshrc", Detail: "→ x"}}},
		{Module: "brew"},
	})
	want := "links:\n  + ~/.zshrc (→ x)\nbrew: up to date\n"
	if err != nil || out.String() != want {
		t.Errorf("got:\n%s\nwant:\n%s", out.String(), want)
	}
}

func TestPlanCount(t *testing.T) {
	p := Plan{{Changes: []Change{{}, {}}}, {}, {Changes: []Change{{}}}}
	if p.Count() != 3 || p.Empty() {
		t.Errorf("Count = %d, Empty = %v", p.Count(), p.Empty())
	}
	if !(Plan{{Module: "x"}}).Empty() {
		t.Error("plan without changes should be empty")
	}
}
