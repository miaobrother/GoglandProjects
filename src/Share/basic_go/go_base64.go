package main

import (
	"fmt"
	"encoding/base64"
)

func main() {
	hello := "你好,世界 hello world"

	 // base64encode
	debyte := base64.StdEncoding.EncodeToString([]byte(hello))
	fmt.Println(debyte)
	//base64decode

	enbyte,err := base64.StdEncoding.DecodeString(string(debyte))

	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(string(enbyte))
}