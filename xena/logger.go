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
}

// newLogger default constructor
func newLogger(debug bool) *logger {
	return &logger{
		debug: debug,
		logger: &logWrapper{
			logger: log.New(os.Stderr, "", log.LstdFlags),
		},
	}
}

// newLogger default constructor
func newEmptyLogger() *logger {
	return &logger{
		logger: &logEmptyWrapper{
			logger: log.New(os.Stderr, "", log.LstdFlags),
		},
	}
}

type logWrapper struct {
	logger *log.Logger
}

// Debugf print debug message. Arguments are handled in the manner of fmt.Printf.
func (l *logWrapper) Debugf(format string, v ...interface{}) {
	l.logger.Printf("Debug: "+format, v...)
}

// Errorf print error message. Arguments are handled in the manner of fmt.Printf.
func (l *logWrapper) Errorf(format string, v ...interface{}) {
	l.logger.Printf("Error: "+format+"\n", v...)
}

type logEmptyWrapper struct {
	logger *log.Logger
}

// Debugf print debug message. Arguments are handled in the manner of fmt.Printf.
func (l *logEmptyWrapper) Debugf(format string, v ...interface{}) {
}

// Errorf print error message. Arguments are handled in the manner of fmt.Printf.
func (l *logEmptyWrapper) Errorf(format string, v ...interface{}) {
}

// logger default logger
type logger struct {
	debug  bool
	logger Logger
}

// SetDebug enable/disable debug logs
func (l *logger) SetDebug(flag bool) {
	l.debug = flag
}

// Debugf print debug message. Arguments are handled in the manner of fmt.Printf.
func (l *logger) Debugf(format string, v ...interface{}) {
	if l.debug {
		l.logger.Debugf(format, v...)
	}
}

// Errorf print error message. Arguments are handled in the manner of fmt.Printf.
func (l *logger) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}
