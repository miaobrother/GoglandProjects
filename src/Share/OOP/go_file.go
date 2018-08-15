package main

import (
	"fmt"
	"os"
)

func WriteFile(path string) {
	//打开文件
	f, err := os.Create("nimei")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i)
		f.WriteString(buf)
		//fmt.Println(n)
	}
}

func ReadFile(path string) {
	//打开文件
	f, err := os.Open("nimei")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	buf := make([]byte, 1024*2)
	n, errOne := f.Read(buf)
	if errOne != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf[:n]))
}

func main() {
	//WriteFile(".")
	ReadFile("nimei")
}
