package iniconfig

import (
	"testing"
	"fmt"
	"io/ioutil"
)

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf MysqlConfig	`ini:"mysql"`
}

type ServerConfig struct {
	Ip string `ini:"ip"`
	Port int	`ini:"port"`
}

type MysqlConfig struct {
	Username string `ini:"username"`
	Passwd string	`ini:"passwd"`
	Database string	`ini:"database"`
	Host string		 `ini:"host"`
	Port int		`ini:"port"`
}



func TestIniConfig(t *testing.T) {
	fmt.Println("hello")
	data,err := ioutil.ReadFile("./config.ini")
	if err != nil{
		fmt.Print("ioutil.ReadFile err is ",err)
		return
	}
	var conf Config
	err = UnMarshal(data,&conf)
	if err != nil{
		t.Error("unmarshal failed, err:",err)
	}
	t.Logf("unmarshal success,conf:%#v\n",conf)
}