package core

import (
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

// UTCFormatter structure for logging
type UTCFormatter struct {
	log.Formatter
}

// Format handler for UTC time - usage: log.SetFormatter(UTCFormatter{&log.JSONFormatter{}})
func (u UTCFormatter) Format(e *log.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	return u.Formatter.Format(e)
}

// init initializes the logger
func init() {
	// Log as JSON instead of the default ASCII formatter.
	//logger.SetFormatter(&logrus.JSONFormatter{})
	log.SetFormatter(UTCFormatter{
		Formatter: &log.TextFormatter{
			DisableColors:   false,
			TimestampFormat: time.RFC3339,
			FullTimestamp:   true,
		},
	})

	// Only log the warning severity or above.
	// Default log level
	log.SetLevel(log.DebugLevel)

	EnvLogLevel := os.Getenv("LOG_LEVEL")
	if EnvLogLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	} else if EnvLogLevel == "info" {
		log.SetLevel(log.InfoLevel)
	} else if EnvLogLevel == "warn" {
		log.SetLevel(log.WarnLevel)
	}
}
