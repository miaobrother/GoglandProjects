package logger

const(
	LogLevelDebug = iota
	LogLevelTrace
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal

)

func getLevelText(level int) string  {
	switch level {
	case LogLevelDebug:
		return "Debug"
	case LogLevelTrace:
		return "Trace"
	case LogLevelInfo:
		return "Info"
	case LogLevelError:
		return "Error"
	case LogLevelWarn:
		return "Warn"
	case LogLevelFatal:
		return "Fatal"
	default:
		return "Unkown"
	}
}