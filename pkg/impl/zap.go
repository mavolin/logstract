package impl

import (
	"go.uber.org/zap"

	"github.com/mavolin/logstract/pkg/logstract"
)

// Zap creates a wrapper for the passed zap SugaredLogger.
func Zap(l *zap.SugaredLogger) logstract.LogFunc {
	return func(lvl logstract.Lvl, msg string, fields logstract.Fields) {
		s := make([]interface{}, 0)

		if fields != nil {
			s = fieldsToSlice(fields)
		}

		switch lvl {
		case logstract.LvlDebug:
			l.Debugw(msg, s)
		case logstract.LvlInfo:
			l.Infow(msg, s)
		case logstract.LvlWarn:
			l.Warnw(msg, s)
		case logstract.LvlError:
			l.Errorw(msg, s)
		case logstract.LvlFatal:
			l.Panicw(msg, s)
		}
	}
}

// fieldsToSlice converts the passed Fields to a slice, as used by zap.
func fieldsToSlice(f logstract.Fields) (s []interface{}) {
	s = make([]interface{}, 2*len(f))

	for k, v := range f {
		s = append(s, k, v)
	}

	return
}
