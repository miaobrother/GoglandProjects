package main

import (
	//"police/download_image"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func initDb() error {
	var err error
	dsn := "police:police123@tcp(test.usnoon.com:3306)/policedata"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}

type CarLog struct {
	Id        int64  `json:"pkid"` //使用tag 对应数据库字段
	LogId     int64  `json:"logid"`
	ImageName string `json:"imagename"`
	Type      int    `json:"type"`
	ListAll   []string
}

func testQueryAllData() (m map[int64]interface{}) { //获取批量数据
	m = make(map[int64]interface{}, 10)
	sqlStr := "select pkid,logid,imagename,type from cmm_uploadlog limit  ?,?"
	rows, err := DB.Query(sqlStr, 0, 10)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}() //设置为rows 必须执行匿名函数；多条记录查询
	if err != nil {
		log.Fatal("Get data list failed: ", err)
	}
	var carlog CarLog
	for rows.Next() { //bool
		err := rows.Scan(&carlog.Id, &carlog.LogId, &carlog.ImageName, &carlog.Type)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return nil
		}
		m[carlog.Id] = carlog.ImageName
	}
	return m
}

func testQueryData() { //获取单挑数据
	sqlStr := "select * from cmm_uploadlog where pkid = ?"
	row := DB.QueryRow(sqlStr, 10)
	var carlog CarLog
	err := row.Scan(&carlog.Id, &carlog.LogId, &carlog.ImageName, &carlog.Type)
	if err != nil {
		log.Fatal("get a data failed:\n", err)
		return
	}

}
func getMaxId() int64 { //放在定时任务里，每三十分钟执行一次
	var carlog CarLog
	sqlStr := "select max(pkid) from cmm_uploadlog"
	res := DB.QueryRow(sqlStr)
	res.Scan(&carlog.Id)
	return carlog.Id
}

func SubCurrenMaxId() int64 {
	s, err := testInsertData()
	g := getMaxId()
	if err != nil {
		log.Fatal("get error failed:", err)
	}
	n := s - g
	fmt.Println("sub a id is:", n)
	return n

}

func testInsertData() (i int64, err error) {
	sqlStr := "insert into cmm_uploadlog(logid,imagename,type)VALUES (?,?,?)"
	res, err := DB.Exec(sqlStr, 49808106, "4b9352b2-857e1323_8_冀A211KK_1506308240098.jpg", 1)
	if err != nil {
		log.Fatal("Insert into cmm_upload failed:", err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal("get lastId failed,", err)
		return
	}
	return lastId, nil
}
func main() {
	err := initDb()
	if err != nil {
		log.Fatal("The initDb failed...")
	}
	//M := testQueryAllData() // 以列表形式获取线上图片
	//download_image.GetImageList(S)
	//for k, v := range M {
	//	fmt.Println(k, v)
	//}
	//testInsertData()//插入数据库cmm_uploadlog数据
	//getMaxId()
	SubCurrenMaxId()
}
