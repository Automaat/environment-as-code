// Package defaults converges macOS preferences via the `defaults` CLI.
package defaults

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Automaat/environment-as-code/internal/config"
	"github.com/Automaat/environment-as-code/internal/engine"
	"github.com/Automaat/environment-as-code/internal/runner"
)

// Processes that cache these domains and only pick up changes on restart.
var restartByDomain = map[string]string{
	"com.apple.dock":          "Dock",
	"com.apple.finder":        "Finder",
	"com.apple.screencapture": "SystemUIServer",
}

type Module struct {
	Defaults []config.Default
	Runner   runner.Runner
}

func (m *Module) Name() string { return "defaults" }

func (m *Module) Plan(ctx context.Context) ([]engine.Change, error) {
	var changes []engine.Change
	var restarts []string
	seen := map[string]bool{}
	for _, d := range m.Defaults {
		want, err := Encode(d.Value)
		if err != nil {
			return nil, fmt.Errorf("%s %s: %w", d.Domain, d.Key, err)
		}
		have, ok, err := m.read(ctx, d.Domain, d.Key)
		if err != nil {
			return nil, err
		}
		if ok && have == want {
			continue
		}
		from := "unset"
		if ok {
			from = have.String()
		}
		changes = append(changes, engine.Change{
			Action: engine.Update,
			Target: d.Domain + " " + d.Key,
			Detail: fmt.Sprintf("%s → %s", from, want),
			Apply:  m.write(d.Domain, d.Key, want),
		})
		if p := restartFor(d); p != "" && !seen[p] {
			seen[p] = true
			restarts = append(restarts, p)
		}
	}
	for _, p := range restarts {
		changes = append(changes, engine.Change{Action: engine.Run, Target: "killall " + p, Apply: m.restart(p)})
	}
	return changes, nil
}

// restart ignores killall's exit code: a process that isn't running reads
// the new value on its next launch anyway.
func (m *Module) restart(process string) func(context.Context) error {
	return func(ctx context.Context) error {
		_, err := m.Runner.Run(ctx, runner.Cmd{Name: "killall", Args: []string{process}})
		return err
	}
}

func restartFor(d config.Default) string {
	if d.Restart != "" {
		return d.Restart
	}
	return restartByDomain[d.Domain]
}

// Value is a typed defaults value in the CLI's textual form.
type Value struct {
	Type string
	Text string
}

func (v Value) String() string { return v.Type + ":" + v.Text }

// Encode converts a YAML value into the form `defaults read` prints, so reads
// and desired values compare directly.
func Encode(v any) (Value, error) {
	switch x := v.(type) {
	case bool:
		if x {
			return Value{"bool", "1"}, nil
		}
		return Value{"bool", "0"}, nil
	case int:
		return Value{"int", strconv.Itoa(x)}, nil
	case float64:
		return Value{"float", strconv.FormatFloat(x, 'f', -1, 64)}, nil
	case string:
		return Value{"string", x}, nil
	}
	return Value{}, fmt.Errorf("unsupported value %#v", v)
}

var typeNames = map[string]string{
	"boolean": "bool",
	"integer": "int",
	"float":   "float",
	"string":  "string",
}

// read returns the current value; ok is false when the key is unset.
// Values of types we don't manage (arrays, dicts, data) are reported with
// their raw type so they always differ from the desired value.
func (m *Module) read(ctx context.Context, domain, key string) (Value, bool, error) {
	res, err := m.Runner.Run(ctx, runner.Cmd{Name: "defaults", Args: []string{"read-type", domain, key}})
	if err != nil {
		return Value{}, false, err
	}
	if res.ExitCode != 0 {
		return Value{}, false, nil
	}
	rawType := strings.TrimPrefix(strings.TrimSpace(res.Stdout), "Type is ")
	typ, known := typeNames[rawType]
	if !known {
		return Value{Type: rawType}, true, nil
	}
	out, err := runner.Output(ctx, m.Runner, runner.Cmd{Name: "defaults", Args: []string{"read", domain, key}})
	if err != nil {
		return Value{}, false, err
	}
	text := strings.TrimSuffix(out, "\n")
	if typ == "float" {
		if f, err := strconv.ParseFloat(text, 64); err == nil {
			text = strconv.FormatFloat(f, 'f', -1, 64)
		}
	}
	return Value{Type: typ, Text: text}, true, nil
}

func (m *Module) write(domain, key string, v Value) func(context.Context) error {
	text := v.Text
	if v.Type == "bool" {
		text = map[string]string{"1": "true", "0": "false"}[text]
	}
	cmd := runner.Cmd{Name: "defaults", Args: []string{"write", domain, key, "-" + v.Type, text}}
	return func(ctx context.Context) error { return runner.Check(ctx, m.Runner, cmd) }
}
