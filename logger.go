package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// Init configures a global Logrus logger.
func Init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

// Info logs at Info level.
func Info(args ...interface{}) {
	log.Info(args...)
}

// Error logs at Error level.
func Error(args ...interface{}) {
	log.Error(args...)
}

// WithField creates a log entry with a custom field.
func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}