package main

import (
	"fmt"
	"math"
)
type circle struct {//定义一个圆的半径结构体
	radius float64
}

type square struct {//定义一个正方形的结构体
	side float64
}

type shape interface {//定义一个统计面积的接口 ，接口调用area函数
	area() float64
}

func (c circle) area() float64  {//执行求圆面积的函数并且返回
	return math.Pi * c.radius * c.radius
}

func ( s square) area() float64  {
	return s.side * s.side
}

func info(z shape)   {
	fmt.Println("This is a info :")
	fmt.Println(z)
	fmt.Println(z.area())
}

func totalArea(shapes... shape) float64  {
	var area float64
	for _,s :=  range shapes{
		area += s.area()
		//fmt.Println(area)
	}
	fmt.Println("以下是shapes:")
	fmt.Println(shapes)
	return area
}

func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
	fmt.Println("This is a Tatol")

	fmt.Println("Total Area:",totalArea(c,s))
}