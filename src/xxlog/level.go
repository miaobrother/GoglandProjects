package xxlog

const (
	XLogLevelDebug = iota
	XLogLevelTrace
	XLogLevelInfo
	XLogLevelWarring
	XLogLevelError
	XLogLevelFatal
)

const (
	XLogTypeFile = iota
	XLogTypeConsole
)
