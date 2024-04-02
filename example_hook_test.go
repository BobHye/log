// +build !windows

package log_test

import (
	"log/syslog"
	"os"

	"github.com/sirupsen/log"
	slhooks "github.com/sirupsen/log/hooks/syslog"
)

// An example on how to use a hook
func Example_hook() {
	var log = log.New()
	log.Formatter = new(log.TextFormatter)                     // default
	log.Formatter.(*log.TextFormatter).DisableColors = true    // remove colors
	log.Formatter.(*log.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	if sl, err := slhooks.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, ""); err == nil {
		log.Hooks.Add(sl)
	}
	log.Out = os.Stdout

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Error("The ice breaks!")

	// Output:
	// level=info msg="A group of walrus emerges from the ocean" animal=walrus size=10
	// level=warning msg="The group's number increased tremendously!" number=122 omg=true
	// level=error msg="The ice breaks!" number=100 omg=true
}
