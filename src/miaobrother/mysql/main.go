package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	//"net/mail"
	"time"
	//"github.com/stackimpact/stackimpact-go"
	"math/rand"
	//"net/url"
)

type UserInfo struct { //将查询到的数据存储到该结构体中
	Userid     int    `Db:"user_id"` //使用tag标示在数据库中字段是什么格式的
	Username   string `Db:"username"`
	Sex        string `Db:"sex"`
	Email      string `Db:"email"`
	Createtime string `Db:"createtime"`
	Age        int    `Db:"age"`
}

func MysqlInsert() {
	Db, err := sqlx.Open("mysql", "root:1q2w3e4r,.@tcp(47.94.197.173:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	//var cost int
	starttime := time.Now().Unix()
	times := int64(time.Now().Nanosecond())
	rand.Seed(times)
	username := fmt.Sprintf("user%d", rand.Int63()) //名字使用随机数
	//fmt.Println(starttime)
	for i := 0; i < 500; i++ {
		_, err := Db.Exec("INSERT INTO person(username,sex,email,createtime,age)VALUES (?,?,?,"+"now()"+",?)", username, "男", "zihao121908@mail.com", 19) //这是插入一条数据  user01,男,zihao121908@mail.com,2018-03-28 10:53:22 ,19
		if err != nil {
			log.Fatal(err)
			return
		}
		/*
			userid ,_ := res.LastInsertId()//返回最后一个 插入的自增id
			fmt.Println("This lastId is",userid)
			fmt.Println(res.LastInsertId())
		*/
	}
	Db.Close()

	endtime := time.Now().Unix()
	fmt.Printf("The pro cost %d s\n", (endtime - starttime))

}

func MysqlSelect() {
	Db, err := sqlx.Open("mysql", "root:1q2w3e4r,.@tcp(47.94.197.173:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(Db)
	fmt.Println()
	starttime := time.Now().UnixNano()
	var userinfo UserInfo // 这是查询单行数据的 必须使用 Get
	err = Db.Get(&userinfo, "select username,sex,email,createtime,age from person where user_id = ?", 432)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("User_info is %#v\n", userinfo)
	fmt.Println("以上是但行数据查询出的结果")
	var userInfoList []*UserInfo //这是查询多行数据的 必须用Select
	err = Db.Select(&userInfoList, "select username,sex,email,createtime,age from person where user_id = ?", 432)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("User_list_info is %#v\n", userInfoList)
	endtime := time.Now().UnixNano()
	fmt.Printf("The pro cost %d s\n", (endtime - starttime))

}
func main() {
	MysqlInsert()
	//MysqlSelect()

}
