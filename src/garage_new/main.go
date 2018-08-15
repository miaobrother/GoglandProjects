package main

import (
	"fmt"
	_ "garage_new/routers"
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin/json"
)

func main() {

	mysqlHost := beego.AppConfig.String("mysqlhost") //配置项获取
	mysqlPort, err := beego.AppConfig.Int("mysqlport")
	if err != nil {
		fmt.Println(err)
	}
	Host := beego.AppConfig.String("host::host") // 获取带有节的数据项配置
	fmt.Println("带有节的主机", Host)
	fmt.Println("The host is", mysqlHost)
	fmt.Println("The host is", mysqlPort)
	m := make(map[string]interface{})
	m["filename"] = "./logs/test.log"
	config, _ := json.Marshal(m)
	beego.SetLogger("file", string(config))
	beego.Debug("service init succ")
	beego.BeeLogger.DelLogger("console")          //只打印到日志文件中
	beego.BConfig.WebConfig.DirectoryIndex = true //该功能类似与ftp工具，访问某一静态页面时可以 列出该目录下所有文件并支持下载
	beego.Run()
}
