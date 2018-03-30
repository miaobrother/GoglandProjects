package main

import "fmt"

type Animal interface {
	Talk()
	Eat()
	Name() string
}

type Dog struct {

}

func (d Dog) Talk()  {
	fmt.Println("汪汪汪")
}
func (d Dog) Eat()  {
	fmt.Println("我在吃骨头...")
}

func (d Dog) Name()(string)  {
	fmt.Println("my name is Zhang")
	return "Zhangsan"
}
func main() {
	var d Dog//定义一个实例
	var a Animal //也是一个实例

	a = d
	a.Eat()
	a.Talk()
	a.Name()
}

