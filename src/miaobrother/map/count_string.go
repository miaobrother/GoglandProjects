package main

import (
	"fmt"
	"strings"
	//"os"
)

func count_Str(str string)map[string]int  {
	var res map[string]int = make(map[string]int)
	words := strings.Split(str," ")
	for _,v := range words{
		fmt.Printf("The v is %v\n",v)
		_,ok:= res[v]
		if !ok{
			res[v] = 1
		}else {
			res[v] ++
		}
	}
	return res
}
func main() {
	var str = "how do you do"
	res := count_Str(str)
	fmt.Printf("The res is %v\n",res)
}
