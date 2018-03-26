package main

import (
	"fmt"
	"os"
	//"image/draw"
	"log"
	"io"
)

func main() {
	CopyFile("target.txt","source.txt")
	fmt.Println("Copy done..")
}

func CopyFile(dstName,srcName string)(written int64,err error)  {
	src,err := os.Open(srcName)//以只读方式打开，并且写入到dst中
	if err != nil{
		log.Fatal(err)
		return
	}
	defer src.Close()
	dst,err := os.OpenFile(dstName,os.O_WRONLY|os.O_CREATE,06666)
	if err != nil{
		return
	}
	defer dst.Close()
	return io.Copy(dst,src)
}
