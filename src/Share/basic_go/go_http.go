package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"go-ethereum/core/types"
	"bytes"
)

func main() {
	//go 发送http请求
	/*
	response,err := http.Get("http://www.baidu.com")

	if err != nil{
		fmt.Println("发送get请求有错误")
		return
	}

	defer response.Body.Close()//类似于析构函数，最后关闭该body请求

	body,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	*/

	//http post 请求
	body := "{\"action\":20}"
	res,err := http.Post("https://passport.baidu.com/v2/api/?login","application/json:charset=utf-8",bytes.NewBuffer([]byte(body)))
	if err != nil{
		fmt.Println("Fatal err",err)
	}

	defer res.Body.Close()

	content,err := ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println("Fatal err",err)
	}
	fmt.Println(string(content))
}