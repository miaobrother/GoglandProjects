package main

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
	"nimeilog/tailf"
)

var (
	appConfig *Config
)

type Config struct {
	logLevel string
	logPath  string
	chanSize int
	kafka    string
	collConf []tailf.CollConf
	etcdAddr string
	etcdKey string
}

func loadCollectConf(conf config.Configer) (err error) {
	var cc tailf.CollConf
	cc.LogPath = conf.String("collect::log_path")
	fmt.Println(cc.LogPath)
	if len(cc.LogPath) == 0 {
		fmt.Println("got a error")
	}
	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("get topic  failed.")
		return

	}

	appConfig.collConf = append(appConfig.collConf, cc)
	return
}

func LoadConf(confType, filename string) (err error) { //加载配置文件中server孑
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed,err:", err)
		return
	}

	appConfig = &Config{}
	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}

	appConfig.logPath = conf.String("logs::log_config")
	if len(appConfig.logPath) == 0 {
		fmt.Println("got log_config failed.")
	}

	appConfig.chanSize, err = conf.Int("logs::chan_size")
	if err != nil {
		appConfig.chanSize = 100
	}

	appConfig.kafka = conf.String("kafka::ip_port")
	if len(appConfig.kafka) == 0 {
		fmt.Println("wuxiao de kafka peizhi")
	}

	appConfig.etcdAddr = conf.String("etcd::addr")
	if len(appConfig.etcdAddr) == 0 {
		fmt.Println("wuxiao de etcd peizhi")
	}

	appConfig.etcdKey = conf.String("etcd::configKey")
	if len(appConfig.etcdKey) == 0 {
		fmt.Println("wuxiao de etcdkey peizhi")
	}

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect conf failed,err is: %v\n", err)
		return
	}
	return

}
