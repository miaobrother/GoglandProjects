package main

import (
	"fmt"
	//"github.com/urfave/cli"
	"time"
)

func main() {
	name := "Naveen"
	fmt.Printf("Orignal String is :%v\n",string(name))
	fmt.Printf("Recversed String:")
	for _,v := range []rune(name) {
		defer fmt.Printf("%c",v)
		time.Sleep(time.Millisecond * 10000)
	}

}