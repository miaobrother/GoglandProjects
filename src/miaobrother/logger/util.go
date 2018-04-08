package logger

import (
	"fmt"
	//"os"
	"path"
	"runtime"
	"time"
)

type logData struct {//存储
	Message  string
	Time     string
	LevelStr string
	Filename string
	FuncName string
	LineNo   int
	WarnAndFatal bool
}

func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}

/*
1、当业务系统调用日志的方法时，可以把日志相关数据写入到chan中
2、只是针对文件日志库
*/

func writeLog(level int, format string, args ...interface{}) *logData {
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(level)

	fileName, funcName, lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format, args...)
	logData := &logData{
		Message:  msg,
		Time:     nowStr,
		LevelStr: levelStr,
		Filename: fileName,
		FuncName: funcName,
		LineNo:   lineNo,
		WarnAndFatal: false,
	}
	if level == LogLevelError || level == LogLevelWarn || level == LogLevelFatal{
		logData.WarnAndFatal = true
	}
	return logData
	//fmt.Fprintf(file,"%s  %s %s:%d %s %s\n",nowStr,levelStr,fileName,lineNo,funcName,msg)
}
