package logstract

import "fmt"

// Logger is the logger used for all logging.
var Logger LogFunc = func(_ Lvl, _ string, _ Fields) {}

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

// Debug creates a debug-level log entry, using the passed message.
func Debug(a ...interface{}) {
	Logger(LvlDebug, fmt.Sprint(a...), nil)
}

// Debugf creates a debug-level log entry, by formatting the passed message.
func Debugf(f string, a ...interface{}) {
	Logger(LvlDebug, fmt.Sprintf(f, a...), nil)
}

// Info creates a info-level log entry, using the passed message.
func Info(a ...interface{}) {
	Logger(LvlInfo, fmt.Sprint(a...), nil)
}

// Infof creates a info-level log entry, by formatting the passed message.
func Infof(f string, a ...interface{}) {
	Logger(LvlInfo, fmt.Sprintf(f, a...), nil)
}

// Warn creates a warning-level log entry, using the passed message.
func Warn(a ...interface{}) {
	Logger(LvlWarn, fmt.Sprint(a...), nil)
}

// Warnf creates a warning-level log entry, by formatting the passed message.
func Warnf(f string, a ...interface{}) {
	Logger(LvlWarn, fmt.Sprintf(f, a...), nil)
}

// Error creates a error-level log entry, using the passed message.
func Error(a ...interface{}) {
	Logger(LvlError, fmt.Sprint(a...), nil)
}

// Errorf creates a error-level log entry, by formatting the passed message.
func Errorf(f string, a ...interface{}) {
	Logger(LvlError, fmt.Sprintf(f, a...), nil)
}

// Fatal creates a fatal-level log entry, using the passed message.
func Fatal(a ...interface{}) {
	Logger(LvlFatal, fmt.Sprint(a...), nil)
}

// Fatalf creates a fatal-level log entry, by formatting the passed message.
func Fatalf(f string, a ...interface{}) {
	Logger(LvlFatal, fmt.Sprintf(f, a...), nil)
}

// Entry is a log entry.
type Entry struct {
	// fields contains the fields of the message
	fields Fields
}

// WithField add the passed field to the entry.
func (e *Entry) WithField(k string, v interface{}) *Entry {
	e.fields[k] = v

	return e
}

// WithFields creates an Entry with the passed Fields.
func (e *Entry) WithFields(f Fields) *Entry {
	for k, v := range f {
		e.fields[k] = v
	}

	return e
}

// Debug creates a debug-level log entry, using the passed message.
func (e *Entry) Debug(a ...interface{}) {
	Logger(LvlDebug, fmt.Sprint(a...), e.fields)
}

// Debugf creates a debug-level log entry, by formatting the passed message.
func (e *Entry) Debugf(f string, a ...interface{}) {
	Logger(LvlDebug, fmt.Sprintf(f, a...), e.fields)
}

// Info creates a info-level log entry, using the passed message.
func (e *Entry) Info(a ...interface{}) {
	Logger(LvlInfo, fmt.Sprint(a...), e.fields)
}

// Infof creates a info-level log entry, by formatting the passed message.
func (e *Entry) Infof(f string, a ...interface{}) {
	Logger(LvlInfo, fmt.Sprintf(f, a...), e.fields)
}

// Warn creates a warning-level log entry, using the passed message.
func (e *Entry) Warn(a ...interface{}) {
	Logger(LvlWarn, fmt.Sprint(a...), e.fields)
}

// Warnf creates a warning-level log entry, by formatting the passed message.
func (e *Entry) Warnf(f string, a ...interface{}) {
	Logger(LvlWarn, fmt.Sprintf(f, a...), e.fields)
}

// Error creates a error-level log entry, using the passed message.
func (e *Entry) Error(a ...interface{}) {
	Logger(LvlError, fmt.Sprint(a...), e.fields)
}

// Errorf creates a error-level log entry, by formatting the passed message.
func (e *Entry) Errorf(f string, a ...interface{}) {
	Logger(LvlError, fmt.Sprintf(f, a...), e.fields)
}

// Fatal creates a fatal-level log entry, using the passed message.
func (e *Entry) Fatal(a ...interface{}) {
	Logger(LvlFatal, fmt.Sprint(a...), e.fields)
}

// Fatalf creates a fatal-level log entry, by formatting the passed message.
func (e *Entry) Fatalf(f string, a ...interface{}) {
	Logger(LvlFatal, fmt.Sprintf(f, a...), e.fields)
}
