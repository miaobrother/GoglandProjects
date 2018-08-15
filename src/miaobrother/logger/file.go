package logger

import (
	"fmt"
	"os"
	//"log"
	//"time"
	//"path"
	//"miaobrother/log"
	"strconv"
	"time"
)

type FileLogger struct { //struct是值类型
	level         int
	logPath       string
	logName       string
	gessfile      *os.File      //定义一个存正常交易日志的文件
	warnFile      *os.File      //定义一个存错误日志的日志文件
	logDataChan   chan *logData //定义一个队列
	logSplitType  int
	logSplitSize  int64
	lastSplitHour int
}

func NewFileLogger(config map[string]string) (log LogInterface, err error) {
	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("not found log_path")
		return
	}
	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("not found log_name")
		return
	}
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not found log_level")
		return
	}

	logChanSize, ok := config["log_chan_size"]
	if !ok {
		logChanSize = "10000"
	}
	chanSize, err := strconv.Atoi(logChanSize)
	if err != nil {
		chanSize = 50000
	}

	var logSplitType int = LogSplitTypeHour
	var logSplitSize int64
	logSplitStr, ok := config["log_split_type"]

	if !ok {
		logSplitStr = "Hour"
	} else {
		if logSplitStr == "size" {
			logSplitSizeStr, ok := config["log_split_size"]
			if !ok {
				logSplitSizeStr = "1048857600"

			}
			logSplitSize, err = strconv.ParseInt(logSplitSizeStr, 10, 64)
			if err != nil {
				logSplitSize = 1048857600
			}
			logSplitType = LogSplitTypeSize
		} else {
			logSplitType = LogSplitTypeSize
		}
	}

	level := getLevelInt(logLevel)
	log = &FileLogger{
		level:         level,
		logPath:       logPath,
		logName:       logName,
		logDataChan:   make(chan *logData, chanSize),
		logSplitSize:  logSplitSize,
		logSplitType:  logSplitType,
		lastSplitHour: time.Now().Hour(),
	}

	fmt.Println(logSplitSize, logSplitType, logSplitStr)
	log.Init() //执行初始化函数
	return

}

func (f *FileLogger) Init() {
	fileName := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("Open file %s failed,err is:%s", fileName, err))
	}
	f.gessfile = file
	//写错误日志fatal & error

	fileName = fmt.Sprintf("%s/%s+warn.log", f.logPath, f.logName)
	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("Open file %s failed,err is:%s", fileName, err))
	}
	f.warnFile = file

	go f.WriteLogBackground() //
}

func (f *FileLogger) splitFileSize(WarnAndFatal bool) {
	file := f.warnFile
	if WarnAndFatal {
		file = f.warnFile
	}
	statInfo, err := file.Stat()
	if err != nil {
		return
	}
	fileSize := statInfo.Size()
	if fileSize <= f.logSplitSize {
		return
	}
	var backupFilename string
	var filename string
	now := time.Now()
	if WarnAndFatal {
		backupFilename = fmt.Sprintf("%s/%s.log.wf%04d%02d%02d%02d%02d%02d", f.logPath, f.logName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)

	} else {
		backupFilename = fmt.Sprintf("%s/%s.log.%04d%02d%02d%02d%02d%02d", f.logPath, f.logName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		filename = fmt.Sprintf("%s/%s.log", f.logPath, f.logName)

	}

	file.Close()
	os.Rename(filename, backupFilename) //备份完成
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return
	}
	if WarnAndFatal {
		f.warnFile = file
	} else {
		f.gessfile = file
	}

}
func (f *FileLogger) splitFileHour(WarnAndFatal bool) {
	now := time.Now()
	hour := now.Hour()
	if hour == f.lastSplitHour {
		return
	}
	var backupFilename string
	var filename string
	if WarnAndFatal {
		backupFilename = fmt.Sprintf("%s/%s.log.wf %04d%02d%02d%02d", f.logPath, f.logName, now.Month(), now.Day(), f.lastSplitHour)
		filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)

	} else {
		backupFilename = fmt.Sprintf("%s/%s.log.%04d%02d%02d%02d", f.logPath, f.logName, now.Year(), now.Month(), now.Day(), f.lastSplitHour)
		filename = fmt.Sprintf("%s/%s.log", f.logPath, f.logName)

	}
	file := f.warnFile
	if WarnAndFatal {
		file = f.warnFile
	}
	file.Close()
	os.Rename(filename, backupFilename) //备份完成
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return
	}
	if WarnAndFatal {
		f.warnFile = file
	} else {
		f.gessfile = file
	}

}

func (f *FileLogger) checkSplitFile(WarnAndFatal bool) {
	if f.logSplitType == LogSplitTypeHour {
		f.splitFileHour(WarnAndFatal)
		return
	}
	f.splitFileSize(WarnAndFatal)

}

func (f *FileLogger) WriteLogBackground() {
	for logData := range f.logDataChan { //每个logData 是一个结构体

		var file *os.File = f.gessfile
		if logData.WarnAndFatal {
			file = f.warnFile
		}
		fmt.Fprintf(file, "%s  %s %s:%d %s %s\n", logData.Time, logData.LevelStr, logData.Filename, logData.LineNo, logData.FuncName, logData.Message)
	}
}

func (f *FileLogger) SetLevel(level int) {
	if f.level <= LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}

	logData := writeLog(LogLevelDebug, format, args...)
	select {
	case f.logDataChan <- logData: //检查管道是否满了，程序不会阻塞
	default: //如果满 了 就走default分支

	}
	//fmt.Fprintf(f.gessfile,  nowStr)
	//fmt.Fprintf(f.gessfile,format,args...)

}
func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	logData := writeLog(LogLevelTrace, format, args...)
	select {
	case f.logDataChan <- logData: //检查管道是否满了，程序不会阻塞
	default:

	}

}
func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	logData := writeLog(LogLevelWarn, format, args...)
	select {
	case f.logDataChan <- logData: //检查管道是否满了，程序不会阻塞
	default:

	}

}
func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	logData := writeLog(LogLevelInfo, format, args...)
	select {
	case f.logDataChan <- logData: //检查管道是否满了，程序不会阻塞
	default:
	}
}
func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	logData := writeLog(LogLevelError, format, args...)
	select {
	case f.logDataChan <- logData: //检查管道是否满了，程序不会阻塞
	default:

	}
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	logData := writeLog(LogLevelFatal, format, args...)
	select {
	case f.logDataChan <- logData: //检查管道是否满了，程序不会阻塞
	default:

	}
}

func (f *FileLogger) Close() {
	f.gessfile.Close()
	f.warnFile.Close()
}
