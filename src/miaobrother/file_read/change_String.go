package main

import (
	"fmt"
)

func main() {
	var s string = "abcdef"
	for i,v := range s{
		fmt.Printf("%d,%c\n",i,v)
	}

	var strToByte []byte
	strToByte = []byte(s)
	strToByte[0] = 'i'//[97 98 99 100 101 102]
//[105 98 99 100 101 102]
	fmt.Println(strToByte)



	var a rune = '磊'
	fmt.Printf(" a is :%c\n",a)

	//var c rune = '你'
	var str = "中文123"

	var runeSlice []rune
	runeSlice = []rune(str)
	fmt.Printf("str len is %d\n",len(runeSlice))
}
