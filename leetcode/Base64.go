package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	msg := "ZmFsY29AdmlwLjE2My5jb20="
	decode, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println(string(decode))
}
