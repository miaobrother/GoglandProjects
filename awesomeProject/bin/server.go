package bin

import (
	"github.com/qiniu/log"
	"github.com/libvirt/libvirt-go"
	//"fmt"
	//"time"
)



func ServerMigrate(source_ip, dest_ip,domName string,source_conn ,dest_conn libvirt.Connect )   {
	//taskresult := TaskResult{}
	//
	//c := &taskresult
	//
	////var taskresult *TaskResult
	//now := time.Now()
	//c.NowTime = now.Format("2006-01-02 03:04:05")

	req := &TaskBody{
		Source_ip: source_ip,
		Dest_ip: dest_ip,
		Domain: domName,
		resultChan: make(chan *TaskResult, 1),
	}

	err := asyncdata.AddRequest(req)

	if err != nil {
		log.Error("add request failed, err:%v", err)
		//c.Result = fmt.Sprintf("add request failed, err:%v", err)
		return
	}
	//
	//
	//
	//resultmap := asyncdata.Migrate(source_conn ,dest_conn,source_ip,domName)
	//
	//c.Result = resultmap
	//
	//defer func() {
	//	end := time.Now().Format("2006-01-02 03:04:05")
	//	c.EndTime =end
	//	elapsed := time.Since(now)
	//	c.RunTime = elapsed
	//}()


}
