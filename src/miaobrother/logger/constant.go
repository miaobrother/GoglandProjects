package logger

const (
	LogLevelDebug = iota
	LogLevelTrace
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)
const (
	LogSplitTypeHour = iota
	LogSplitTypeSize
)

func getLevelText(level int) string {
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
		return "Debug"
	}
}

func getLevelInt(level string) int {
	switch level {
	case "debug":
		return LogLevelDebug
	case "trace":
		return LogLevelTrace
	case "info":
		return LogLevelInfo
	case "error":
		return LogLevelError
	case "warn":
		return LogLevelWarn
	case "fatal":
		return LogLevelFatal
	}
	return LogLevelDebug
}
