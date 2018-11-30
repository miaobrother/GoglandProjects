package xxlog

import (
	"fmt"
)

type Xlog interface {
	LogDebug(fmt string, args ...interface{})
	LogTrace(fmt string, args ...interface{})
	LogInfo(fmt string, args ...interface{})
	LogWarring(fmt string, args ...interface{})
	LogError(fmt string, args ...interface{})
	LogFatal(fmt string,args ...interface{})

	SetLevel(level int)
}

func NewXLog(logType,level int,filename string) Xlog  {
	var logger Xlog
	switch (logType) {
	case XLogTypeFile:
		logger = &XFile{
			level:level,
			filename:filename,
		}

	case XLogTypeConsole:
		logger = &XConsole{
			level:level,
	}
	default:
		logger = &XFile{
			level:level,
			filename:filename,
		}

	}
}
