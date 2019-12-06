package xena

import (
	"log"
	"os"
)

// Logger interface
type Logger interface {
	// Debugf print debug message. Arguments are handled in the manner of fmt.Printf.
	Debugf(format string, v ...interface{})
	// Errorf print error message. Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, v ...interface{})
	// // SetDebug enable/disable debug logs
	SetDebug(flag bool)
}

// newLogger default constructor
func newLogger() Logger {
	return &logger{
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

// logger default logger
type logger struct {
	debug  bool
	logger *log.Logger
}

// SetDebug enable/disable debug logs
func (l *logger) SetDebug(flag bool) {
	l.debug = flag
}

// Debugf print debug message. Arguments are handled in the manner of fmt.Printf.
func (l *logger) Debugf(format string, v ...interface{}) {
	if l.debug {
		l.logger.Printf("Debug: "+format, v...)
	}
}

// Errorf print error message. Arguments are handled in the manner of fmt.Printf.
func (l *logger) Errorf(format string, v ...interface{}) {
	l.logger.Printf("Error: "+format, v...)
}
