package bin

import (
	"github.com/libvirt/libvirt-go"
	"fmt"

	"sync"

	"time"
	"net"
	"awesomeProject/lib"
	"github.com/jmoiron/sqlx"

	"encoding/xml"
	"crypto/md5"

	"encoding/hex"
	"log"
)

var ConnMap map[string]libvirt.Connect


var DenyMap map[string]string


type ConnMain struct {
	ConnChan chan libvirt.Connect
	ServerList []string
	msgChan chan string
	newClientChan chan net.Conn
	lock sync.RWMutex


}
type LibvrtdDB struct {
	ServerIp string `db:"server_ip"`
	ID int `db:"id"`

}

type Result struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type GetDominInfo struct{
	Code int
	Message interface{}
}

func NewMgr() *ConnMain  {
	Mgr := &ConnMain{
		ServerList:make([]string,10),
		ConnChan:make(chan libvirt.Connect,1),

	}
	ConnMap = make(map[string]libvirt.Connect,1024)
	DenyMap = make(map[string]string,1024)
	//go Mgr.StartProcess()

	return Mgr
}

func (L *ConnMain) Select() (libvrtdDBs []*LibvrtdDB,err error) {

	Db,err :=sqlx.Open("mysql","cmdb_u:k4RmQxzTM@tcp(10.1.8.120:3306)/cmdb_db")
	if err !=nil{
		log.Println("connect to mysql failed ",err)
		return
	}

	var libvrtdDB []*LibvrtdDB
	err =Db.Select(&libvrtdDB,"SELECT server_ip,id from vmanageplatform_vmserver;",)
	if err!=nil{
		log.Printf("SELECT table error",err)
		return
	}
	defer Db.Close()

	return libvrtdDB,err
}
//定时更新ConnMap最新连接
func (L *ConnMain)Run(){
	t := time.NewTicker(time.Second *3)
	for range t.C{

		go func() {

			libvrtdDBs,err:= L.Select()
			if err!=nil{
				log.Println("Select ServerIp false")
			}
			for i:=0;i<len(libvrtdDBs);i++{
				L.PortMinion(libvrtdDBs[i].ServerIp)
			}
		}()
	}
	defer t.Stop()
}



func (L *ConnMain) PortMinion(server_ip string) (err error){
	defer func() {
		err :=recover()
		if err !=nil{
			//fmt.Printf("PortMinion ip%s,recover,error:%s\n",server_ip,err)
			return
		}

	}()

	url :=fmt.Sprintf("%s:16509",server_ip)
	net_conn, err := net.DialTimeout("tcp", url, time.Millisecond * 59)
	if err!=nil{
		L.lock.RLock()
		_,exists :=ConnMap[server_ip]
		L.lock.RUnlock()
		if exists{
			L.lock.Lock()
			delete(ConnMap,server_ip)
			DenyMap[server_ip] = "Flase"
			L.lock.Unlock()
			log.Printf("删除:%s",server_ip)

		}
	}

	L.lock.RLock()
	_,exist := ConnMap[server_ip]
	L.lock.RUnlock()

	if exist == false  {
		//log.Println("定时扫描 结果", exist,server_ip)
		L.Process(server_ip)
	}
	net_conn.Close()
	return
}
//获取所有Conn连接
//func (L *ConnMain) PUSHAllConn(server_id int,server_ip string){
//	err:=L.PortMinion(server_ip)
//	if err!=nil{
//		return
//	}
//	L.Process(server_ip)
//
//}
//
//func (L *ConnMain) StartProcess() {
//	libvrtdDBs,err:= L.Select()
//	if err!=nil{
//		fmt.Println("Select ServerIp false")
//	}
//	for i:=0;i<len(libvrtdDBs);i++{
//		L.PUSHAllConn(libvrtdDBs[i].ID,libvrtdDBs[i].ServerIp)
//	}
//}
func (L *ConnMain) Process(server_ip string){
	L.lock.RLock()
	_,exist := ConnMap[server_ip]
	L.lock.RUnlock()
	if exist {
		log.Println("已经存在字典",exist)
		return
	}else {
		url := fmt.Sprintf("qemu+tcp://%s/system",server_ip)
		libvirt_conn,err:=libvirt.NewConnect(url)
		if err!=nil{
			L.lock.Lock()
			delete(ConnMap,server_ip)
			DenyMap[server_ip] = "Flase"
			L.lock.Unlock()
			//fmt.Printf("Error libvirt_conn%s,%s",libvirt_conn,server_ip)
			return
		}
		L.ConnChan <- *libvirt_conn
		select {
		case <-L.ConnChan:
			L.lock.Lock()
			ConnMap[server_ip] = *libvirt_conn
			L.lock.Unlock()
			log.Println("更新到字典中",server_ip)

		}


		//L.GetAllInfo(server_ip)
		//defer libvirt_conn.Close()
	}


}
//获取所有信息
func (L *ConnMain) GetAllInfo(server_ip string){
	fmt.Println("GetAllInfo")
	for ids,conn :=range ConnMap{
		fmt.Println("conn",ids)
		L.GetDomainINfo(&conn,&ids)
	}
}

//获取单台宿主机所有信息
func (L *ConnMain) GetDomainINfo(conn *libvirt.Connect,server_ip *string) (callback GetDominInfo) {
	//fmt.Printf("GetDomainINfo server_ip%s \n",*server_ip)
	doms_run_active, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	doms_run_shutdown,err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_SHUTOFF)
	lib.Eroor(err)
	doms_list := make([]libvirt.Domain,0)
	doms_list = append(doms_list,doms_run_active...)
	doms_list = append(doms_list,doms_run_shutdown...)
	fmt.Println("doms_run_active",len(doms_run_active))
	fmt.Println("doms_run_paused",len(doms_run_shutdown))
	json_data :=make([]interface{},0)
	callback =GetDominInfo{}
	for _,v :=range doms_list{
		name,_ :=v.GetName()
		ids,_ := v.GetID()
		info_dict := make(map[string]interface{},0)
		info_dict["name"] = name

		vmserverinfo,_:=lib.SelectVmServerPersonInfo(&name)
		info_dict["ids"] = ids
		//查看设置最大内存
		//_, err = v.GetMaxMemory()
		//CPU个数
		info,_:=v.GetVcpus()
		//运行状态
		DomainState, _,_ :=v.GetState()
		info_dict["status"] = DomainState
		var vcpu_number int
		for range info{
			vcpu_number +=1
		}
		info_dict["vcpu"] = vcpu_number
		//fmt.Println("v.Free()",v.GetState())
		//conn.G
		//info_dict["mem_per"] = v.GetMemoryParameters()
		//fmt.Printf("name:%s id:%d\n",name ,ids,stats)
		instance,_ := v.GetXMLDesc(0)
		data :=[]byte(instance)
		v := lib.DomainXml{}
		err = xml.Unmarshal(data,&v)


		//if err!=nil{
		//	fmt.Println("Unmarshal ERROR",err)
		//}
		vnc :=v.RDevices.Graphics.Port
		dev :=v.RDevices.Interface.Target.Dev
		info_dict["vnc"] = vnc
		info_dict["dev"] =dev

		//fmt.Println("LOOK:",name,stats,err,dev,vnc)
		testdata := fmt.Sprintf("%s%s",*server_ip,name)
		t := md5.New()
		t.Write([]byte(testdata))
		md5Data := t.Sum([]byte(""))
		info_dict["token"] = hex.EncodeToString(md5Data)


		for i:=0;i<len(vmserverinfo);i++{
			info_dict["expire_time"] = vmserverinfo[i].Expire_time.String
			info_dict["item_name"] = vmserverinfo[i].Item_name.String
			info_dict["item_person"] = vmserverinfo[i].Item_person.String
			info_dict["item_person_email"] = vmserverinfo[i].Item_person_email.String
			info_dict["ip"] = vmserverinfo[i].Belong_to_groups.String

		}
		json_data = append(json_data, info_dict)

		}
	callback.Code = 200
	callback.Message = json_data

	return callback

}


func (L *ConnMain) Start(conn *libvirt.Connect,domain_name string) (result Result)  {
	result = Result{}

	//result := make(map[string]string,0)
	domain,err:=conn.LookupDomainByName(domain_name)
	if err!=nil{
		result.Code = 1006
		result.Message = err.Error()

		return result
	}
	err=domain.Create()
	if err!=nil{
		result.Code = 1007
		result.Message = err.Error()
		//return result
		return result
	}
	result.Code = 200
	result.Message = "Success"
	fmt.Println("Start",domain)
	return result

}


func (L *ConnMain) Shutdown(conn *libvirt.Connect,domain_name string) (result Result){
	result = Result{}

	//result := make(map[string]string,0)
	domain,err:=conn.LookupDomainByName(domain_name)
	if err!=nil{
		result.Code = 1006
		result.Message = err.Error()

		return result
	}
	err=domain.Shutdown()
	if err!=nil{
		result.Code = 1007
		//libvirt.Error.Message
		result.Message = err.Error()
		return result
	}
	result.Code = 200
	result.Message = "Success"
	fmt.Println("Start",domain)

	return result
}





