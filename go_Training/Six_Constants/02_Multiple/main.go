package main

import (
	"fmt"
)

//const (
//	pi = 3.14
//	language = "Go"
//)
//
//func main() {
//	fmt.Println(pi)
//	fmt.Println(language)
//}

//const (
//	a = iota //0
//	b = iota * 10//1
//	c = iota * 10 //2
//)
//
//const (
//	d = iota //0
//	e //1
//	f  //2
//)
//
//func main() {
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(c)
//	fmt.Println(d)
//	fmt.Println(e)
//	fmt.Println(f)
//}

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // * 1的 10次方
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)

}
