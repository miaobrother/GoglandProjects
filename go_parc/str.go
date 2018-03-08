package main

import (
	"fmt"
)
func sliceToMap(s []string)map[string]int  {
	sli := make(map[string]int)
	for _,v := range s{
		_,exist := sli[v]
		if exist{
			sli[v] += 1
		}else {
			sli[v] =1
		}

	}
	return sli

}
func main() {
	//test123()
	var strs = []string{"lisi","zhangsan","Nimei","张三","lisi","张三"}
	a := sliceToMap(strs)
	//sliceToMap()
	fmt.Println(a)


}



