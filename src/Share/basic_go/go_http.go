package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"go-ethereum/core/types"
	//"bytes"
	"net/url"
	"strings"
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

	/*
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
	*/

	//设置url请求的参数
	v := url.Values{}
	v.Set("mobile","13733208110")
	//body
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))
	client := &http.Client{}
	//此项是赋值的请求：需要在请求时设置头参数、cookie 之类的数据

	request,err := http.NewRequest("POST","https://passport.baidu.com/?getpassusertype&tt=1504531745839",body)

	if err != nil{
		fmt.Println("Fatal err",err.Error())
	}

	request.Header.Set("Content-Type","application/x-www-form-urlencoded;param=value")

	resp ,err := client.Do(request)
	defer resp.Body.Close()
	content,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("Fatal error",err.Error())
	}
	fmt.Println(string(content))
}