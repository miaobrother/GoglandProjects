package main

import (
	"Share/iniconfig"
	"fmt"
)
type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf  MysqlConfig  `ini:"mysql"`
}

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type MysqlConfig struct {
	Username string `ini:"username"`
	Passwd   string `ini:"passwd"`
	Database string `ini:"database"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Price    float32 `ini:"price"`
}


func main() {
	filename := "/Users/zhangyalei/test.conf"
	var conf Config
	err := iniconfig.UnMarshalFile(filename,&conf)
	if err != nil{
		fmt.Println("unmarshal failed,err: ",err)
		return
	}
	fmt.Printf("conf: %#v\n",conf)

}