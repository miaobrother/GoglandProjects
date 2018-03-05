package main

import (
	"fmt"
)

func main() {
	nums := []int{1,2,3,4,5}
	sum := 0

	for _,num := range nums{
		sum += num
	}

	fmt.Println("sum is:",sum)

	kvs := map[string]string{"Name":"zhangsan","Sex":"F"}
	for k,v := range kvs{
		fmt.Printf("%v : %v\n",k,v)
	}


	for i ,c := range "go"{
		fmt.Println(string(i),string(c))
	}


}