package log_test

import (
	"os"

	"github.com/sirupsen/log"
)

var (
	mystring string
)

type GlobalHook struct {
}

func (h *GlobalHook) Levels() []log.Level {
	return log.AllLevels
}

func (h *GlobalHook) Fire(e *log.Entry) error {
	e.Data["mystring"] = mystring
	return nil
}

func ExampleGlobalHook() {
	l := log.New()
	l.Out = os.Stdout
	l.Formatter = &log.TextFormatter{DisableTimestamp: true, DisableColors: true}
	l.AddHook(&GlobalHook{})
	mystring = "first value"
	l.Info("first log")
	mystring = "another value"
	l.Info("second log")
	// Output:
	// level=info msg="first log" mystring="first value"
	// level=info msg="second log" mystring="another value"
}
