package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}//定义一个基本的表示几何图形的接口

//正方形及圆形都实现这个接口
type square struct {
	width,height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64  {
	return s.width * s.height
}

func (s square) perim() float64  {
	return 2*s.width + 2*s.height
}//正方形面积

func (c circle) area() float64  {
	return math.Pi*c.radius*c.radius
}

func (c circle) perim() float64  {
	return 2 * math.Pi *c.radius
}

func measure(g geometry)  {
	fmt.Println("The g is %v",g)

	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	s := square{width:3,height:4}

	c:= circle{radius:5}

	measure(s)
	measure(c)
}


