package database

import (
	"log"
)

// GooseLogger adapts logging for Goose migrations
type GooseLogger struct {
	Verbose bool
}

// Printf implements goose.Logger
func (l *GooseLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Fatalf implements goose.Logger
func (l *GooseLogger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
