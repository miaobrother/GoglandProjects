package main

import (
	"fmt"
)

func testfanzhuanV2(s []rune) string  {
	var result string
	for i := 0;i < len(s);i++{
		result = result + fmt.Sprintf("%c",s[len(s)-i-1])
	}
	return result
}
func testfanzhuanstr(s string) string  {
	var byteS = []rune(s)
	for from, to := 0,len(byteS)-1; from < to; from,to = from +1,to -1{
		byteS[from],byteS[to] = byteS[to],byteS[from]
	}
	return string(byteS)
}

func huiwen() string {
	var s = "上海自来水来自海上"
	var r = []rune(s)
	for from, to := 0,len(r)-1; from < to; from,to = from +1,to -1{
		r[from],r[to] = r[to],r[from]
	}
	if string(r) == s{
		fmt.Println("The s is huiwen ")
	}else {
		fmt.Println("The s isn't huiwen ")
	}
	return string(r)


}


func main() {
	var s = "woai ni 爱"
	fmt.Println(testfanzhuanstr(s))
	ss := []rune(s)
	fmt.Println(testfanzhuanV2(ss))

	huiwen()
}

