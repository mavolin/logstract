package logstract

type Lvl uint8

const (
	// LvlDebug is the debug level.
	LvlDebug Lvl = iota
	// LvlInfo is the info level.
	LvlInfo
	// LvlWarn is the warning level.
	LvlWarn
	// LvlError is the error level.
	LvlError
	// LvlFatal is the fatal level.
	// This level preferably causes a panic or if not supported by the logger a normal os.Exit.
	LvlFatal
)

type (
	// Fields is a map of fields used for the log entry.
	Fields map[string]interface{}

	// LogFunc is the abstraction of the logger.
	LogFunc func(lvl Lvl, msg string, fields Fields)
)
