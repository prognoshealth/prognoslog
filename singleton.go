package prognoslog

import (
	"os"
	"sync"
)

var singletonLog *Logger // nolint
var once sync.Once       // nolint

// SingletonLog exposes a single log instance that can be used within an
// application that writes to stdout with a verbosity of false
func SingletonLog() *Logger {
	once.Do(func() {
		singletonLog = &Logger{out: os.Stdout, isVerbose: false}
	})
	return singletonLog
}
