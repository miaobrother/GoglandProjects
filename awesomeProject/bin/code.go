package bin

import "fmt"

const (
	Success = 200
	LibvirtdError = 500
	MysqlConnError = 1001
	ActionawargsError = 1002
	LibvritdServerIPError = 1003
	LibvirtdConnError = 1004
	ActionkwargsError = 1005
	LibvirtdLookupDomainByNameError = 1006
	LibvirtdRunActionError = 1007


)
func GetMessgae(code int,values interface{})(msg string){
	switch code {
	case LibvirtdError:
		msg = ""
	case Success:
		msg = "操作成功"
	case ActionawargsError:

		msg = fmt.Sprintf("参数:%s不合法",values)
	case LibvritdServerIPError:
		msg = fmt.Sprintf("数据库中无此ID:%d",values)
	case LibvirtdConnError:
		msg = fmt.Sprintf("此IP:%s未建立连接",values)
	case ActionkwargsError:
		msg = fmt.Sprintf("方法%s不存在",values)
	case MysqlConnError:
		msg = fmt.Sprintf("Mysql连接失败")
	case LibvirtdRunActionError:

		msg = fmt.Sprintf("获取域:%s失败",values)
	case LibvirtdLookupDomainByNameError:
		msg = fmt.Sprintf("操作失败",values)

	}
	return msg
}