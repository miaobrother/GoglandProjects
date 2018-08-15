package main

import (
	"fmt"
	"time"
)

func test() {
	now := time.Now()
	fmt.Println(now.Day())
	fmt.Println(now.Unix())
	fmt.Println(now.Nanosecond())
}

func testTimestamp(timestamp int64) {
	timeObj := time.Unix(time.Now().Unix(), 0)
	year := timeObj.Year()
	fmt.Println(year)
}
func processTask() {
	fmt.Printf("go to task\n")
}
func testTicker() { //时间管理器
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Printf("%v\n", i)
		processTask()
	}

}

func timeFmt() {
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05") //按照格式化内容
	fmt.Println(timeStr)
}

func testCost() { //计算程序耗时多长时间
	start := time.Now().UnixNano() //1970 年到此时的纳秒数
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
	}
	end := time.Now().UnixNano()
	cost := (end - start) / 1000
	fmt.Printf("code cost:%d\n", cost)
}

func main() {
	//test()
	//timestamp := time.Now().Unix()
	//testTimestamp(timestamp)
	//testTicker()
	//timeFmt()
	testCost()
}
