package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	//获取当前时间

	now := time.Now()
	p(now) //2018-03-07 14:51:58.463526664 +0800 CST

	then := time.Date(2018,3,6,19,20,23,987654321,time.UTC)
	p(then) //2018-03-06 19:20:23.987654321 +0000 UTC


	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}