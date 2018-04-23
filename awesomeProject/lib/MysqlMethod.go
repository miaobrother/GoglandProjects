package lib

import (

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
)
type VMserverDB struct {
	ServerIp string `db:"server_ip"`
	ID int `db:"id"`

}

type VmServerPersonInfo struct {
	Expire_time sql.NullString `db:"expire_time"`
	Item_name sql.NullString `db:"item_name"`
	Belong_to_groups sql.NullString `db:"belong_to_groups"`
	Item_person sql.NullString`db:"item_person"`
	Item_person_email sql.NullString `db:"item_person_email"`

}
func MysqlConn() (Dbs sqlx.DB,err error){

	Db,err :=sqlx.Open("mysql","cmdb_u:k4RmQxzTM@tcp(10.1.8.120:3306)/cmdb_db")
	return *Db,err
}

func SelectVmServer() (vmservers []*VMserverDB,err error) {
	Db,err :=MysqlConn()
	if err !=nil{
		fmt.Println("connect to mysql failed ",err)
	}
	var vmserver []*VMserverDB
	err =Db.Select(&vmserver,"SELECT server_ip,id from vmanageplatform_vmserver;",)
	if err!=nil{
		fmt.Printf("SELECT table error",err)
	}
	Db.Close()
	return vmserver,err
}


func SelectVmServerPersonInfo(servername *string)(vmserverinfos []*VmServerPersonInfo,err error){
	Db,err :=MysqlConn()
	if err !=nil{
		fmt.Println("connect to mysql failed ",err)
	}
	var vmserverinfo []*VmServerPersonInfo
	err =Db.Select(&vmserverinfo,"SELECT expire_time,item_name,belong_to_groups,item_person, item_person_email from vmanageplatform_vmserverinstance WHERE name=?",servername)
	if err!=nil{
		fmt.Printf("SELECT table error",err)
	}
	Db.Close()

	return vmserverinfo,err


}