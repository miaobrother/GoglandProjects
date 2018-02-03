package main

import (
	"fmt"
	"os"
	"flag"
)

var (
	conf  string
	level int
)

func init() {
	flag.StringVar(&conf, "c", "/Users/zhangyalei/Desktop/text.txt", "请")
	flag.IntVar(&level, "1", 8, " 请指定日志级别")
}

func main() {

	fmt.Printf("conf is %s\n", conf)
	fmt.Printf("level is :%d\n", level)
	for index, val := range os.Args {
		fmt.Printf("args[%d] = %s\n", index, val)
	}
}
