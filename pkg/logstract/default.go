package logstract

import "fmt"

// DefaultLogger is the global logger.
var DefaultLogger *Logger

// Init initializes the global logger.
func Init(l LogFunc) {
	DefaultLogger = NewLogger(l)
}

// WithField creates an Entry with the passed field.
func WithField(k string, v interface{}) *Entry {
	return &Entry{
		fields: Fields{
			k: v,
		},
	}
}

// WithFields creates an Entry with the passed Fields.
func WithFields(f Fields) *Entry {
	return &Entry{
		fields: f,
	}
}

// Debug creates a debug-level log entry, using the passed message and the DefaultLogger.
func Debug(a ...interface{}) { DefaultLogger.l(LvlDebug, fmt.Sprint(a), nil) }

// Debugf creates a debug-level log entry, by formatting the passed message and the DefaultLogger.
func Debugf(f string, a ...interface{}) { DefaultLogger.l(LvlDebug, fmt.Sprintf(f, a), nil) }

// Info creates a info-level log entry, using the passed message and the DefaultLogger.
func Info(a ...interface{}) { DefaultLogger.l(LvlInfo, fmt.Sprint(a), nil) }

// Infof creates a info-level log entry, by formatting the passed message and the DefaultLogger.
func Infof(f string, a ...interface{}) { DefaultLogger.l(LvlInfo, fmt.Sprintf(f, a), nil) }

// Warn creates a warning-level log entry, using the passed message and the DefaultLogger.
func Warn(a ...interface{}) { DefaultLogger.l(LvlWarn, fmt.Sprint(a), nil) }

// Warnf creates a warning-level log entry, by formatting the passed message and the DefaultLogger.
func Warnf(f string, a ...interface{}) { DefaultLogger.l(LvlWarn, fmt.Sprintf(f, a), nil) }

// Error creates a error-level log entry, using the passed message and the DefaultLogger.
func Error(a ...interface{}) { DefaultLogger.l(LvlError, fmt.Sprint(a), nil) }

// Errorf creates a error-level log entry, by formatting the passed message and the DefaultLogger.
func Errorf(f string, a ...interface{}) { DefaultLogger.l(LvlError, fmt.Sprintf(f, a), nil) }

// Fatal creates a fatal-level log entry, using the passed message and the DefaultLogger.
func Fatal(a ...interface{}) { DefaultLogger.l(LvlFatal, fmt.Sprint(a), nil) }

// Fatalf creates a fatal-level log entry, by formatting the passed message and the DefaultLogger.
func Fatalf(f string, a ...interface{}) { DefaultLogger.l(LvlFatal, fmt.Sprintf(f, a), nil) }
