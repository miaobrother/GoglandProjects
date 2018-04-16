package main

import (
	"fmt"
	//"github.com/astaxie/beego"
)

func main() {
	//如果超过容量 则按照两倍扩容
	s := make([]int,1,2)
	oldCap := cap(s)
	//fmt.Printf("The old cap is :%d\n",oldCap)

	for i := 0;i < 8;i ++{
		s = append(s,i)
		if newCap := cap(s);oldCap< newCap{
			//oldCap = newCap
			fmt.Printf("cap :%d ---> %d\n",oldCap,newCap)
			oldCap = newCap
		}
	}
	//fmt.Printf("The old cap is :%d\n",oldCap)

}