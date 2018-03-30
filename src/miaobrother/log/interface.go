package log

type LogInterface interface {
	logDebug(msg string)
	logWarn(msg string)
}
