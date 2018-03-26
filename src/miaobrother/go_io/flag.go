package main

import (
	"flag"
	"fmt"
)

var recu bool
var test string
var level int

func init() {
	flag.BoolVar(&recu,"r",false,"use recu..")
	flag.StringVar(&test,"t","defalut value..","string option")
	flag.IntVar(&level,"1",1,"level is a value")
	flag.Parse()
}

func main() {
	fmt.Printf("recu is %v\n",recu)
	fmt.Printf("test:%v\n",test)
	fmt.Printf("level:%v\n",level)
}