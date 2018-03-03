package main

import "fmt"


func intSeq() func() int{
	i := 0

	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	// 为了确认闭包的状态是独立于intSeq函数的，再创建一个
	newInts := intSeq()
	fmt.Println(newInts()) //  1
}