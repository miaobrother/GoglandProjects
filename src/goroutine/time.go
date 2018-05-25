package main

import (
	"time"
	"fmt"
)


func main() {
	//获取当前时间戳
	now := time.Now().Unix()
	fmt.Println(now)
	fmt.Println(time.Now())

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))


	str_timeOne := time.Unix(0,1).Format("2006-01-02 15:04:05")
	fmt.Println(str_timeOne)

	str_timeTwo := time.Unix(1527231931,0).Format("2006年01月02日 15点04分05秒")
	fmt.Println(str_timeTwo)


	the_time := time.Date(2018, 3, 30, 15, 24, 59, 0,time.Local)
	unix_time := the_time.Unix()
	fmt.Println(unix_time)


	t := time.NewTicker(time.Second*1)

	for v := range t.C{

		fmt.Println("hello",v)
		return
	}

}
