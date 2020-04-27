package impl

import (
	"github.com/sirupsen/logrus"

	"github.com/mavolin/logstract/pkg/logstract"
)

// Logrus creates a wrapper for the passed logrus logger.
func Logrus(l *logrus.Logger) logstract.LogFunc {
	return func(lvl logstract.Lvl, msg string, fields logstract.Fields) {
		e := logrus.NewEntry(l)

		if len(fields) != 0 {
			e.WithFields(logrus.Fields(fields))
		}

		switch lvl {
		case logstract.LvlDebug:
			e.Debug(msg)
		case logstract.LvlInfo:
			e.Info(msg)
		case logstract.LvlWarn:
			e.Warn(msg)
		case logstract.LvlError:
			e.Error(msg)
		case logstract.LvlFatal:
			e.Panic(msg)
		}
	}
}
