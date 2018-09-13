package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func ConvertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "info":
		return logs.LevelInfo
	case "warn":
		return logs.LevelWarn
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

func InitLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = appConfig.logPath
	config["level"] = ConvertLogLevel(appConfig.logLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("InitLooger failed ,err: ", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	return

}
