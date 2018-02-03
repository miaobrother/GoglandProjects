package main

import "fmt"

func test1() {
	var a int
	a = 10
	fmt.Printf("The a is %d\n", a)

	var b *int
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &b)

	b = &a
	fmt.Printf("%d\n", *b)

	fmt.Printf("%p\n", &b)
}

func modify(a *int) {
	*a = 100
}
func test2()  {
	var b = 1
	var p *int
	p = &b
	modify(p)
	fmt.Println(b)
}

func test3()  {
	///P 默认初始化nil
	var p *int
	var b int
	p = &b
	*p = 100
	p = new(int)
	*p = 1000
	fmt.Println(*p)

	//指针类型的变量初始化：1 使用其他变量地址给它赋值， 2 使用new分配
}

func test4()  {
	//First
	var p *string
	p = new(string)

	*p = "abc"
	fmt.Println(*p)

	//Two
	var str string = "hello world"
	p = &str
	fmt.Println(*p)
}

func test5()  {
	var a []int
	a = make([]int,10)
	a[0] = 99
	fmt.Println(a)

	var p *[]int
	p = new([]int)

	(*p) = make([]int,10)
	(*p)[0] = 100

	fmt.Println(p)


}
func modify_arr(a [6] int)  {
	a[0] = 100
}
func test6()  {
	var a[6] int

	modify_arr(a)
	fmt.Println(a)
}

func add(a,b int)int  {
	return a+b
}


func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	c := add
	fmt.Printf("%p %T  %p %T\n",c,add,c,add)

	sum :=c(10,20)
	fmt.Println(sum)
	sum = add(10,20)
	fmt.Println(sum)
}
