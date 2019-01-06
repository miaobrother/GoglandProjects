package main

import "fmt"

func main(){
	//切片与元组的区别 数组的长度是一个const，长度不能修改
	a := []int{1,2,3}
	b := [3]int{1,2,3}
	fmt.Printf("The a type is %T\n",a)
	fmt.Printf("The b type is %T\n",b)
	//fmt.Printf("The a = b is %v\n",a == b)
	c := append(a, 4)
	fmt.Printf("The c len is %d\n The c cap is %d\n",len(c),cap(c))
	s2 := make([]string,0,2)
	fmt.Printf("The s2 len is %d\n The s2 cap is %d\n",len(s2),cap(s2))
}
