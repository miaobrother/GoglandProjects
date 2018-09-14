package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//   数据源名称
	db, err := sql.Open("mysql", "dev:13733208110smC!5%d+gH#sa1+@tcp(43.254.1.150:3306)/go_test?charset=utf8")
	checkError(err)
	//插入数据
	stmt, err := db.Prepare("INSERT test_go set NAME=?,password = ?,other = ?,ctime = ?,utime =?")
	checkError(err)
	//标准格式化时间
	ctime := time.Now().Format("2006-01-02 15:04:05")
	password := []byte("test")
	p := md5.Sum(password)

	//执行插入
	res, err := stmt.Exec("test1", hex.EncodeToString(p[:]), "test", ctime, ctime)
	checkError(err)

	//获取最新插入的insertid

	id, err := res.LastInsertId()
	checkError(err)
	fmt.Println(id)

	//查询数据

	rows, err := db.Query("select * from test_go")
	checkError(err)
	for rows.Next() {
		var id int
		var name string
		var password string
		var other string
		var ctime string
		var utime string
		err = rows.Scan(&id, &name, &password, &other, &ctime, &utime)
		checkError(err)

		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(password)
		fmt.Println(other)
		fmt.Println(ctime)
		fmt.Println(utime)
	}
}
