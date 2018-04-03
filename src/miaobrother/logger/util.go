package logger

import (
	"runtime"
	"os"
	"time"
	"path"
	"fmt"
)

type logData struct {

	Message string
	Time string
	LevelStr string
	Filename string
	FuncName string
	LineNo int
}

func GetLineInfo()(fileName string,funcName string,lineNo int)  {
	pc,file,line,ok := runtime.Caller(3)
	if ok{
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

func writeLog(file *os.File,level int,format string,args...interface{})  {
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(level)

	fileName,funcName,lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format,args...)
	fmt.Fprintf(file,"%s  %s %s:%d %s %s\n",nowStr,levelStr,fileName,lineNo,funcName,msg)
}
