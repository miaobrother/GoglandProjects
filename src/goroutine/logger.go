package main

import (
	"fmt"
	"log"
	"os"
)

func Debug(logName string) {
	logFile, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		fmt.Printf("create test.log err : %v\n", err)
		return
	}
	defer logFile.Close()

	debugLog := log.New(logFile, "[Debug]", log.Ldate)

	debugLog.SetPrefix("[Debug]")
	debugLog.SetFlags(log.Lshortfile)
	debugLog.Println("this is Debug log")
}

func Waring(logName string) {
	logFile, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("create waring.log err: %v\n", err)
	}
	defer logFile.Close()

	waringLog := log.New(logFile, "[Waring]", log.Ldate)

	waringLog.SetPrefix("[Waring]")
	waringLog.SetFlags(log.Lshortfile)
	waringLog.Println("This is a waring log")
}

func main() {
	//logName := "Waring" + "/test.log"
	Debug("/Users/xxx/Debug" + "test.log")
	Waring("/Users/xxx/Waring" + "test.log")
}
