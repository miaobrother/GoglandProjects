package main

import "fmt"

func test1() {
	var a [5]int
	b := a[1:3]
	b[0] = 200
	b[1] = 100
	fmt.Printf("a:%#v\n", a)
}

func test2() {
	var a [5]int
	b := a[1:3]
	a[0] = 100
	a[1] = 200
	fmt.Printf("b :%#v\n", b)
}

func Sum(a [100]int) int {
	var sum int
	for i := 0; i < len(a); i++ {
		sum = sum + a[i]

	}

	return sum
}

func testApp() {
	var a []int = []int{1, 2, 3}
	a = append(a, 1, 2, 3)
	fmt.Printf("a:%#v\n", a)
}

func testStrSlice() {
	var str = "hello world"
	str1 := str[6:]
	fmt.Printf("str1:%s\n", str1)
}

func testChangeStr()  {
	var str = "hello world"
	var b []byte = []byte(str)
	b[0] = 'a'
	str1 := string(b)
	fmt.Printf("%s\n,%d\n",str1,len(str))
}
func main() {
	//test1()
	//test2()
	//var a [5]int
	//b := a[1:3]
	//a[0] =100
	//a[1] = 200
	//fmt.Printf("b :%#v\n",b)
	//var a [100]int =[100]int{1,2,3,4,5,6,7,8,9}
	//res := Sum(a)
	//fmt.Printf("sum = %d\n", res)
	//testApp()
	//testStrSlice()
	testChangeStr()
}
