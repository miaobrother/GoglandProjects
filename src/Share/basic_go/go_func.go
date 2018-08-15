package main

import "fmt"

type Circle struct {
	radius float64
}

func main() {

	var cOne Circle
	cOne.radius = 10.0
	fmt.Println("Area of Circle(cOne=", cOne.getArea())

	fmt.Println("----------")
	//闭包和普通函数的区别
	tmp := []int{1, 2, 3}

	for _, i := range tmp {
		fmt.Println(i)
		test(i)
	}

	fmt.Println("-----------")

	for _, i := range tmp {
		fmt.Println(i)
		defer test(i)
	}

	fmt.Println("-------------")

	for _, i := range tmp {
		fmt.Println(i)
		defer func() {
			fmt.Println(i)
		}()
	}
}

//普通函数
func test(i int) {
	fmt.Println(i)
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
