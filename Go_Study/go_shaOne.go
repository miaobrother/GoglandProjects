package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string" //源字符串

	h := sha1.New()//生成一个hash模式
	h.Write([]byte(s))//必须使用字节数组

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n",bs)
}
