package hooks

import (
	"github.com/sirupsen/logrus"
)

const (
	ServerVersionField = "version"
)

// RequestIDHook is a Logrus hook for including request ID in log entries
type ServerVersionHook struct {
	ServerVersion string
}

// Levels returns the logging levels for which this hook should be called
func (hook *ServerVersionHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire is called when a log entry is made
func (hook *ServerVersionHook) Fire(entry *logrus.Entry) error {
	entry.Data[ServerVersionField] = hook.ServerVersion
	return nil
}
