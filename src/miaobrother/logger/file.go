package logger

import (
	"os"
	"fmt"
	//"log"
	//"time"
	//"path"
	//"miaobrother/log"
)

type FileLogger struct {
	level   int
	logPath string
	logName string
	gessfile *os.File //定义一个存正常交易日志的文件
	warnFile *os.File //定义一个存错误日志的日志文件
}

func NewFileLogger(level int, logPath, logName string) LogInterface {
	logger :=  &FileLogger{
		level:   level,
		logPath: logPath,
		logName: logName,
	}
	logger.init()//执行初始化函数
	return logger

}

func (f *FileLogger) init()  {
	fileName := fmt.Sprintf("%s/%s.log",f.logPath,f.logName)
	file,err := os.OpenFile(fileName,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0755)
	if err != nil{
		panic(fmt.Sprintf("Open file %s failed,err is:%s",fileName,err))
	}
	f.gessfile = file
	//写错误日志fatal & error

	fileName = fmt.Sprintf("%s/%s+warn.log",f.logPath,f.logName)
	file,err = os.OpenFile(fileName,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0755)
	if err != nil{
		panic(fmt.Sprintf("Open file %s failed,err is:%s",fileName,err))
	}
	f.warnFile = file
}




func (f *FileLogger) SetLevel(level int)  {
	if f.level <= LogLevelDebug || level > LogLevelFatal{
		level = LogLevelDebug
	}
	f.level = level
}



func (f *FileLogger) Debug(format string,args...interface{})  {
	if f.level > LogLevelDebug{
		return
	}
	writeLog(f.gessfile,LogLevelDebug,format,args...)
	//fmt.Fprintf(f.gessfile,  nowStr)
	//fmt.Fprintf(f.gessfile,format,args...)

}
func (f *FileLogger) Trace(format string,args...interface{})  {
	if f.level > LogLevelTrace{
		return
	}
	writeLog(f.gessfile,LogLevelTrace,format,args...)

}
func (f *FileLogger) Warn(format string,args...interface{})  {
	if f.level > LogLevelWarn{
		return
	}
	writeLog(f.warnFile,LogLevelWarn,format ,args...)

}
func (f *FileLogger) Info(format string,args...interface{})  {
	if f.level > LogLevelInfo{
		return
	}
	writeLog(f.gessfile,LogLevelInfo,format,args...)
}
func (f *FileLogger) Error(format string,args...interface{})  {
	if f.level > LogLevelError{
		return
	}
	writeLog(f.warnFile,LogLevelError,format,args...)
}

func (f *FileLogger) Fatal(format string,args...interface{})  {
	if f.level > LogLevelFatal{
		return
	}
	writeLog(f.warnFile,LogLevelFatal,format,args...)
}

func (f *FileLogger) Close()  {
	f.gessfile.Close()
	f.warnFile.Close()
}