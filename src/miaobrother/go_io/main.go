package main

import (
	//"bufio"
	"fmt"
	//"io"
	"os"
	"io"
)

func main() {
	inputFile ,err := os.Open("./movie.go")//只读方式打开
	if err != nil{
		fmt.Printf("open file err:%v\n",err)
		return
	}
	var buf[123]byte//设置读取的字节数
	var sli []byte
	//inputFile.Read(buf[:])
	for {//循环读取123字节
		n, err := inputFile.Read(buf[:])
		if err == io.EOF{ //此时代表文件已经读完了
			break
		}else if err != nil{
			fmt.Println("read filed:",err)
			return
		}
		sli = append(sli,buf[:n]...)

	}
	fmt.Println(string(sli))
	defer inputFile.Close()
}
