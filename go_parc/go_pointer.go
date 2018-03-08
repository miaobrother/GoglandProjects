package main

import "fmt"


//我们用两个不同的例子来演示指针的用法

func zeroval(ival int)  {
	ival = 0
}

func zeroptr(iptr *int)  {
	*iptr = 0
}

func main() {
	i := 1

	fmt.Println("initial:",i)

	zeroval(i)
	fmt.Println("zeroval is:",i)

	zeroptr(&i)
	fmt.Println("zeroptr is:",i)
	fmt.Println("pointer:",&i)
}

