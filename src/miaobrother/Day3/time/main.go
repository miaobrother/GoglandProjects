package main

import (
	"fmt"
	"time"
)

func testTime() {
	for {
		now := time.Now()
		fmt.Printf("type oof now is %T\n", now)

		year := now.Year()
		month := now.Month()
		day := now.Day()

		str := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d\n", year, month, day, now.Hour(), now.Minute(), now.Second())
		fmt.Printf("%s", str)
		time.Sleep(time.Second)
	}
}
func testTimeCost() {
	start := time.Now().UnixNano()
	/*
		业务层面code
	*/
	time.Sleep(time.Millisecond*10)
	end := time.Now().UnixNano()

	cost := (end - start) / 1000
	fmt.Printf("The cost is :%dus\n", cost)
}
func main() {

	//testTime()
	testTimeCost()

}
