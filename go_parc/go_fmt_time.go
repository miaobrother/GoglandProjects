package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	t := time.Now()
	p(t)

	p(t.Format("2006-01-02T15:04:05Z07:00"))
	p(t.Format("3:04PM"))//这是一个模版

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d:%d\n",t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second(),t.Nanosecond())
}