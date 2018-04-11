package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"go-ethereum/core/types"
)

func main() {
	//go 发送http请求
	response,err := http.Get("http://www.baidu.com")

	if err != nil{
		fmt.Println("发送get请求有错误")
		return
	}

	defer response.Body.Close()//类似于析构函数，最后关闭该body请求

	body,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}