package main

import (
	//"miaobrother/log"
	"miaobrother/logger"
	//"fmt"
	"time"
)


func initLogger(name,logPath,logName string,level string)(err error)  {
	m := make(map[string]string,8)
	m["log_path"] = logPath
	m["log_name"] = "user_server"
	m["log_level"] = level
	err = logger.InitLogger(name,m)
	if err != nil{
		return
	}
	logger.Debug("init logger success..")
	return
}

func Run()  {
	for{
		logger.Debug("User server is running...")
		time.Sleep(time.Second)
	}

}

func main() {
	/*
	file := log.FileLogOne("")
	file.LogDebug("This is a debug info")
	file.LogWarn("This is a warn info")
	*/

	/*
	console := log.NewConsoleLog("")
	console.LogDebug("This is a console debug info")
	console.LogWarn("This is a console warn info")
	*/

	//log := log.FileLogOne("") // 这是使用的接口 打印到file中
	initLogger("console","/Users/zhangyalei","gesspsbc.log","debug")
	Run()
	return
}