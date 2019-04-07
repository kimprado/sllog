package logging

import (
	"log"
	"strings"
)

// LoggerWriter represents loggerWirter
type LoggerWriter interface {
	write(message string, v ...interface{})
}

// LoggerWriterOut message logging
type LoggerWriterOut struct {
	name  string
	level string
}

// LoggerWriterOut send message to logging
func (l *LoggerWriterOut) write(message string, v ...interface{}) {
	log.Printf(l.level+" "+l.name+" "+message, v...)
}

// LoggerWriterDiscard just discard messages
type LoggerWriterDiscard struct {
}

// Discards messages
func (l *LoggerWriterDiscard) write(message string, v ...interface{}) {
}

// Logger allows logging
type Logger struct {
	name string

	traceLoggger LoggerWriter
	debugLoggger LoggerWriter
	infoLoggger  LoggerWriter
	warnLoggger  LoggerWriter
	errorLoggger LoggerWriter
	fatalLoggger LoggerWriter

	isTraceEnabled bool
	isDebugEnabled bool
	isInfoEnabled  bool
	isWarnEnabled  bool
	isErrorEnabled bool
	isFatalEnabled bool

	level loggerLevel
}

// NewLogger creates new logger
func NewLogger(name string, config map[string]string) Logger {

	loggerLevel := loggerEnabled(name, config)
	level := loggerLevel.level

	logger := Logger{
		name:         name,
		traceLoggger: &LoggerWriterDiscard{},
		debugLoggger: &LoggerWriterDiscard{},
		infoLoggger:  &LoggerWriterDiscard{},
		warnLoggger:  &LoggerWriterDiscard{},
		errorLoggger: &LoggerWriterDiscard{},
		fatalLoggger: &LoggerWriterDiscard{},
		level:        loggerLevel,
	}

	if level == "" {
		return logger
	}

	if level == "TRACE" {
		logger.isTraceEnabled = true
		logger.traceLoggger = &LoggerWriterOut{name: name, level: "TRACE"}
	}

	if level == "DEBUG" || logger.isTraceEnabled {
		logger.isDebugEnabled = true
		logger.debugLoggger = &LoggerWriterOut{name: name, level: "DEBUG"}
	}

	if level == "INFO" || logger.isDebugEnabled {
		logger.isInfoEnabled = true
		logger.infoLoggger = &LoggerWriterOut{name: name, level: "INFO"}
	}

	if level == "WARN" || logger.isInfoEnabled {
		logger.isWarnEnabled = true
		logger.warnLoggger = &LoggerWriterOut{name: name, level: "WARN"}
	}

	if level == "ERROR" || logger.isWarnEnabled {
		logger.isErrorEnabled = true
		logger.errorLoggger = &LoggerWriterOut{name: name, level: "ERROR"}
	}

	if level == "FATAL" || logger.isErrorEnabled {
		logger.isFatalEnabled = true
		logger.fatalLoggger = &LoggerWriterOut{name: name, level: "FATAL"}
	}

	return logger
}

// IsTraceEnabled -
func (l *Logger) IsTraceEnabled() bool {
	return l.isTraceEnabled
}

// IsDebugEnabled -
func (l *Logger) IsDebugEnabled() bool {
	return l.isDebugEnabled
}

// IsInfoEnabled -
func (l *Logger) IsInfoEnabled() bool {
	return l.isInfoEnabled
}

//IsWarnEnabled -
func (l *Logger) IsWarnEnabled() bool {
	return l.isWarnEnabled
}

// IsErrorEnabled -
func (l *Logger) IsErrorEnabled() bool {
	return l.isErrorEnabled
}

// IsFatalEnabled -
func (l *Logger) IsFatalEnabled() bool {
	return l.isFatalEnabled
}

// Fatalf - fatal
func (l *Logger) Fatalf(message string, v ...interface{}) {
	l.fatalLoggger.write(message, v...)
}

// Errorf - erro
func (l *Logger) Errorf(message string, v ...interface{}) {
	l.errorLoggger.write(message, v...)
}

// Warnf - debug
func (l *Logger) Warnf(message string, v ...interface{}) {
	l.warnLoggger.write(message, v...)
}

// Infof - info
func (l *Logger) Infof(message string, v ...interface{}) {
	l.infoLoggger.write(message, v...)
}

// Debugf - debug
func (l *Logger) Debugf(message string, v ...interface{}) {
	l.debugLoggger.write(message, v...)
}

// Tracef - trace
func (l *Logger) Tracef(message string, v ...interface{}) {
	l.traceLoggger.write(message, v...)
}

// getLoggerEnabled returns active LoggerLevel. Uses 'ROOT' as default
func loggerEnabled(name string, config map[string]string) loggerLevel {
	var loggerLevelEnabled loggerLevel

	loggerLevelValue := config[name]
	if loggerLevelValue != "" {
		loggerLevelEnabled = loggerLevel{name, loggerLevelValue}
		return loggerLevelEnabled
	}

	var rootLevel loggerLevel
	for k, v := range config {
		if k == "ROOT" {
			rootLevel = loggerLevel{k, v}
			continue
		}
		if !strings.Contains(name, k) {
			continue
		}
		if len(loggerLevelEnabled.name) > len(k) {
			continue
		}
		loggerLevelEnabled = loggerLevel{k, v}
	}
	if loggerLevelEnabled.level == "" {
		loggerLevelEnabled = rootLevel
	}

	return loggerLevelEnabled
}

type loggerLevel struct {
	name  string
	level string
}
