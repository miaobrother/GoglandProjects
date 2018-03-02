package main

import (
	"net/http"
	//"net"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	res ,err := http.Get("http://192.168.111.111:8111/deviceinfo")
	if err != nil{
		log.Fatal(err)
	}
	defer res.Body.Close()

	body,err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Println(err)
	}
	m := map[string]string{}
	fmt.Println(string(body))
}