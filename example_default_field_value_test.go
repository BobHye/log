package log_test

import (
	"os"

	"github.com/sirupsen/log"
)

type DefaultFieldHook struct {
	GetValue func() string
}

func (h *DefaultFieldHook) Levels() []log.Level {
	return log.AllLevels
}

func (h *DefaultFieldHook) Fire(e *log.Entry) error {
	e.Data["aDefaultField"] = h.GetValue()
	return nil
}

func ExampleDefaultFieldHook() {
	l := log.New()
	l.Out = os.Stdout
	l.Formatter = &log.TextFormatter{DisableTimestamp: true, DisableColors: true}

	l.AddHook(&DefaultFieldHook{GetValue: func() string { return "with its default value" }})
	l.Info("first log")
	// Output:
	// level=info msg="first log" aDefaultField="with its default value"
}
