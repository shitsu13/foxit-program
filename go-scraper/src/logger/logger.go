package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	Writer        io.Writer
	ContextLogger *log.Entry
}

func NewLogger(appName string) *Logger {
	writer := os.Stderr
	contextLogger := log.WithFields(log.Fields{
		"app_name": appName,
	})

	return &Logger{
		Writer:        writer,
		ContextLogger: contextLogger,
	}
}
