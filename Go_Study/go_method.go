package main

import "fmt"

type rect struct {
	width,height int
}

//area 方法有一个限定类型 *rect
//表示这个函数是定义在 rect 结构体上的方法
func (r *rect) area() int  {
	return r.width * r.height
}

func (r rect) perim()  int  {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width:10,height:5}

	fmt.Println("area:",r.area())
	fmt.Println("perim:",r.perim())


	rp := &r

	fmt.Println("area:" ,rp.area())
	fmt.Println("perim:",rp.perim())
}