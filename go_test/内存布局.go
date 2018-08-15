package main

import "fmt"

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
	//r1 内存布局
	fmt.Printf("x addr is %p\n", &r1.p1.x)
	fmt.Printf("y addr is %p\n", &r1.p1.y)
	fmt.Printf("x addr is %p\n", &r1.p2.x)
	fmt.Printf("y addr is %p\n", &r1.p2.y)
	fmt.Printf("r2 x addr is %p\n", &r2.p1.x)
	fmt.Printf("r2 y addr is %p\n", &r2.p1.y)
	fmt.Printf("r2 x addr is %p\n", &r2.p2.x)
	fmt.Printf("r2 y addr is %p\n", &r2.p2.y)

}
