package main

import (
	"math"
	"fmt"
)

type Vertex struct {//定义了一个Vertex 类。
	X,Y float64
}

func (c Vertex) As() float64   {//为类Vertex添加Abs方法，该方法返回一个float64位的值。
	return (math.Sqrt(c.X*c.X + c.Y*c.Y))
}

func (c  *Vertex) Scale(f float64)  {//为类Vertex添加Scale方法，并使用pointer
	c.X = c.X *f
	c.Y = c.Y *f
}

func main() {
	c:=Vertex{3,4}
	c.Scale(10)
	fmt.Println(c.As())
}