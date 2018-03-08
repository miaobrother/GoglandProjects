package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func checkOne(e error)  {
	if e != nil{
		panic(e)
	}
}

func main() {
	dOne := []byte("hello\ngo\n")
	err := ioutil.WriteFile("new.log",dOne,0644)
	checkOne(err)

	f,err := os.Create("ew.log")
	checkOne(err)
	defer f.Close()

	dTwo := []byte{115,111,109,101,10}
	n2,err := f.Write(dTwo)
	checkOne(err)
	fmt.Printf("wrote %d bytes\n",n2)


	n3,err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n",n3)

	f.Sync()//调用Sync方法来将缓冲区数据写入磁盘

	//bufio除了提供上面的缓冲区读取数据外，还提供了缓冲写入数据的方法

	w:= bufio.NewWriter(f)
	n4,err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n",n4)

	w.Flush()
}