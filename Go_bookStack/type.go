package main

import "fmt"

type TZ int
type STR string

func main() {
	var a,b TZ = 3,4
	var m,n STR = "zhangsan","lisi"
	fmt.Println("M+n is:",m+n)
	fmt.Printf("a+b is: %d\n",a+b)
	c := a+b
	fmt.Printf("%d\n",c)
}
