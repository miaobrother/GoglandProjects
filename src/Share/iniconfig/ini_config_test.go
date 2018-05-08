package iniconfig

import (
	"fmt"
	"io/ioutil"
	"testing"
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

func TestIniConfig(t *testing.T) { //单元测试函数
	fmt.Println("hello begin test config.ini file ")
	data, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		fmt.Print("ioutil.ReadFile err is ", err)
		return
	}
	var conf Config
	err = UnMarshal(data, &conf)
	if err != nil {
		t.Error("unmarshal failed, err:", err)
	}
	confData, err := Marshal(conf)
	fmt.Println(string(confData))
	if err != nil{
		fmt.Errorf("the err is %s\n",err)
	}
	t.Logf("unmarshal success,conf:%#v\n", conf)

	MarshalFile(conf,"/Users/zhangyalei/Desktop/test.conf")
}


func TestIniConfigFileOpen(t *testing.T) { //单元测试函数
	filename := "/Users/zhangyalei/test.conf"
	var conf Config
	err := MarshalFile(conf,filename)
	if err != nil{
		t.Errorf("unmarshal failed,err:%v\n",err)
		return
	}
	var confTwo Config
	err = UnMarshalFile(filename,&confTwo)
	//fmt.Println(string(confData))
	if err != nil{
		fmt.Errorf("the err is %s\n",err)
	}
	t.Logf("unmarshal failed,conf:%#v\n", confTwo)
}