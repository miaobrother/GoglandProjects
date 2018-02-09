package main

import "fmt"

func main() {
	i , j := 42,1024
	p := &i// 生成一个指向i的指针，使用*p来打印出i的值。
	fmt.Println(*p,j)// 42 ,1024


	*p = 21
	fmt.Println(i)//i 是21


	p = &j
	fmt.Println(*p) //1024
	*p = *p / 37
	fmt.Printf("The *p is %T\n",*p)
}