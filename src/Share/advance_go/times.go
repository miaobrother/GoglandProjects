package main

import (
	"fmt"
	"time"
	"os"
	//"log"
)

func main() {
	t := time.Now().Format("2006-01-02")
	fmt.Println(t)
	dir := os.Mkdir(t,0777)
	fmt.Println(dir)
}