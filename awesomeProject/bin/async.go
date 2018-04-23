package bin

import (
	"sync"
	"github.com/libvirt/libvirt-go"
	"time"
	"fmt"
)


//一个任务返回的具体信息
type TaskResult struct {
	NowTime string
	StartTime string
	EndTime   string
	RunTime	time.Duration
	Status    int
	Result    interface{}
}

//一个任务
type TaskBody struct {
	Source_ip string
	Dest_ip string
	Domain string

	resultChan chan *TaskResult
}


//所有task信息
type AsyncData struct {


	rwLock sync.RWMutex

	AsyncQueue chan *TaskBody
}

var asyncdata *AsyncData


func (a *AsyncData) AddRequest(req *TaskBody) (err error){

	timer := time.NewTicker(3*time.Second)
	defer func() {
		timer.Stop()
	}()

	select {
	case a.AsyncQueue <- req:
	case <-timer.C:
		err = fmt.Errorf("time out")
		return
	}
	return

}

func (a *AsyncData) Migrate(source_conn ,dest_conn libvirt.Connect,source_ip,domName string) (resultmap map[string]interface{}){

	source_dom,_:=source_conn.LookupDomainByName(domName)
	_,err := source_dom.Migrate(&dest_conn,libvirt.MIGRATE_LIVE,"","",0)

	resultmap = make(map[string]interface{},0)
	if err !=nil{

		resultmap["code"] = 404
		resultmap["msg"]= fmt.Sprintf("源主机%s,迁移domain：%s失败;失败信息:%s",source_ip,domName,err)


		return resultmap

	}


	fmt.Println("source_conn",source_conn,dest_conn)

	resultmap["code"] = 200
	resultmap["msg"]= fmt.Sprintf("主机%s,迁移domain：%s成功,",source_ip,domName)


	return resultmap




}


func InitAsync() *AsyncData {

	req := &AsyncData{
		AsyncQueue:make(chan *TaskBody,1024),
	}

	return req
}