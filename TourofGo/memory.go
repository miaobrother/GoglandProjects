package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func do() {
	max := 2100000
	item := 15111
	var a []int
	for range make([]struct{}, max) {
		a = append(a, item)
	}
	var sitem []string
	for range make([]struct{}, max) {
		sitem = append(sitem, strconv.Itoa(item))
	}
	fmt.Println(len(sitem))
	runtime.GC()
}

func main() {
	for range make([]struct{}, 10) {
		go do()
	}
	time.Sleep(time.Second * 30)
}
