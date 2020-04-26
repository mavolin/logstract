package logstract

import "fmt"

// Logger wraps LogFunc to provide convenient logging.
type Logger struct {
	// l is the LogFunc used for logging.
	l LogFunc
}

// NewLogger creates a new Logger using the passed LogFunc.
func NewLogger(l LogFunc) *Logger {
	return &Logger{
		l: l,
	}
}

// WithField creates an Entry with the passed field.
func (l *Logger) WithField(k string, v interface{}) *Entry {
	return &Entry{
		fields: Fields{
			k: v,
		},
	}
}

// WithFields creates an Entry with the passed Fields.
func (l *Logger) WithFields(f Fields) *Entry {
	return &Entry{
		fields: f,
	}
}

// Debug creates a debug-level log entry, using the passed message.
func (l *Logger) Debug(a ...interface{}) { l.l(LvlDebug, fmt.Sprint(a), nil) }

// Debugf creates a debug-level log entry, by formatting the passed message.
func (l *Logger) Debugf(f string, a ...interface{}) { l.l(LvlDebug, fmt.Sprintf(f, a), nil) }

// Info creates a info-level log entry, using the passed message.
func (l *Logger) Info(a ...interface{}) { l.l(LvlInfo, fmt.Sprint(a), nil) }

// Infof creates a info-level log entry, by formatting the passed message.
func (l *Logger) Infof(f string, a ...interface{}) { l.l(LvlInfo, fmt.Sprintf(f, a), nil) }

// Warn creates a warning-level log entry, using the passed message.
func (l *Logger) Warn(a ...interface{}) { l.l(LvlWarn, fmt.Sprint(a), nil) }

// Warnf creates a warning-level log entry, by formatting the passed message.
func (l *Logger) Warnf(f string, a ...interface{}) { l.l(LvlWarn, fmt.Sprintf(f, a), nil) }

// Error creates a error-level log entry, using the passed message.
func (l *Logger) Error(a ...interface{}) { l.l(LvlError, fmt.Sprint(a), nil) }

// Errorf creates a error-level log entry, by formatting the passed message.
func (l *Logger) Errorf(f string, a ...interface{}) { l.l(LvlError, fmt.Sprintf(f, a), nil) }

// Fatal creates a fatal-level log entry, using the passed message.
func (l *Logger) Fatal(a ...interface{}) { l.l(LvlFatal, fmt.Sprint(a), nil) }

// Fatalf creates a fatal-level log entry, by formatting the passed message.
func (l *Logger) Fatalf(f string, a ...interface{}) { l.l(LvlFatal, fmt.Sprintf(f, a), nil) }

// Entry is a log entry.
type Entry struct {
	// fields contains the fields of the message
	fields Fields
	// l is the LogFunc used for logging.
	l LogFunc
}

// WithField add the passed field to the entry.
func (l *Entry) WithField(k string, v interface{}) *Entry {
	if l.fields == nil {
		l.fields = Fields{
			k: v,
		}
	} else {
		l.fields[k] = v
	}

	return l
}

// WithFields creates an Entry with the passed Fields.
func (l *Entry) WithFields(f Fields) *Entry {
	if l.fields == nil {
		l.fields = f
	} else {
		for k, v := range f {
			l.fields[k] = v
		}
	}

	return l
}

// Debug creates a debug-level log entry, using the passed message.
func (l *Entry) Debug(a ...interface{}) { l.l(LvlDebug, fmt.Sprint(a), nil) }

// Debugf creates a debug-level log entry, by formatting the passed message.
func (l *Entry) Debugf(f string, a ...interface{}) { l.l(LvlDebug, fmt.Sprintf(f, a), nil) }

// Info creates a info-level log entry, using the passed message.
func (l *Entry) Info(a ...interface{}) { l.l(LvlInfo, fmt.Sprint(a), nil) }

// Infof creates a info-level log entry, by formatting the passed message.
func (l *Entry) Infof(f string, a ...interface{}) { l.l(LvlInfo, fmt.Sprintf(f, a), nil) }

// Warn creates a warning-level log entry, using the passed message.
func (l *Entry) Warn(a ...interface{}) { l.l(LvlWarn, fmt.Sprint(a), nil) }

// Warnf creates a warning-level log entry, by formatting the passed message.
func (l *Entry) Warnf(f string, a ...interface{}) { l.l(LvlWarn, fmt.Sprintf(f, a), nil) }

// Error creates a error-level log entry, using the passed message.
func (l *Entry) Error(a ...interface{}) { l.l(LvlError, fmt.Sprint(a), nil) }

// Errorf creates a error-level log entry, by formatting the passed message.
func (l *Entry) Errorf(f string, a ...interface{}) { l.l(LvlError, fmt.Sprintf(f, a), nil) }

// Fatal creates a fatal-level log entry, using the passed message.
func (l *Entry) Fatal(a ...interface{}) { l.l(LvlFatal, fmt.Sprint(a), nil) }

// Fatalf creates a fatal-level log entry, by formatting the passed message.
func (l *Entry) Fatalf(f string, a ...interface{}) { l.l(LvlFatal, fmt.Sprintf(f, a), nil) }
