package main

import (
	"fmt"
	//"go/types"
)

type Point struct {
	x int
	y int
}

type Rect struct {
	p1 Point
	p2 Point
}

type RectA struct {
	p1 *Point
	p2 *Point
}

func main() {
	var r1 Rect
	var r2 RectA
	r2.p1 = new(Point)
	r2.p2 = &Point{}
	//r1的内存布局
	fmt.Printf("p1.x addr is :%p\n", &r1.p1.x)
	fmt.Printf("p1.y addr is :%p\n", &r1.p1.y)
	fmt.Printf("p2.x addr is :%p\n", &r1.p2.x)
	fmt.Printf("p2.y addr is :%p\n", &r1.p2.y)

	//r2的内存布局
	fmt.Printf("p1.x addr is :%p\n", &r2.p1.x)
	fmt.Printf("p1.y addr is :%p\n", &r2.p1.y)
	fmt.Printf("p2.x addr is :%p\n", &r2.p2.x)
	fmt.Printf("p2.y addr is :%p\n", &r2.p2.y)
}
