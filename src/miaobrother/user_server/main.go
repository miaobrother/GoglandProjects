package main

import (
	"miaobrother/log"
	//"fmt"
)

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
	log := log.NewConsoleLog("") //这是使用接口 打印到console中  很简介 逼使用单个的file以及console方便很多
	log.LogDebug("This is a debug info ")
	log.LogWarn("This is a warn info")
}