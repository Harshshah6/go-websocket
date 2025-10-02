package logger

import (
	"log"
	"os"
)

// Logger is a simple wrapper around the standard logger
type Logger struct {
	*log.Logger
}

// New creates a new logger with a given prefix
func New(prefix string) *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, prefix+" ", log.LstdFlags|log.Lshortfile),
	}
}
