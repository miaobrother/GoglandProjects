package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)//编码

	sDec,_:= b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))// 解码
	 fmt.Println()

	 uEnc := b64.URLEncoding.EncodeToString([]byte(data))

	 fmt.Println(uEnc)//编码

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))//兼容模式解码
}