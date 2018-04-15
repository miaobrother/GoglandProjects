package main

import "fmt"

func swap(a,b *int)  {
	*a,*b = *b,*a
	fmt.Printf("in the swap:a = %d b = %d\n",*a,*b)
}

func main() {
	var a,b = 10,20
	swap(&a,&b)//变量本身传递，值传递，并非是地址的传递
	fmt.Printf("in the main:a = %d b = %d\n",a,b)

}