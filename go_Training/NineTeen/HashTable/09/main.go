package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	res,err := http.Get("http://www.baidu.com")

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(res.Body)
	bs,_ := ioutil.ReadAll(res.Body)
	str := string(bs)

	fmt.Println(str)

}