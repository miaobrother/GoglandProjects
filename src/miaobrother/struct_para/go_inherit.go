package main

import (
	"fmt"
	//"expvar"
)

type Animal struct {
	Name string
	Sex string
}

func (a *Animal) Talk()  {
	fmt.Printf("I'm talk,I'am %s\n",a.Name)
}

type Dog struct {
	Feet string
	//Animal//这是值类型 是拷贝
	*Animal //通常是一个指针类型,此时Dog具有Animal的特性
}

func (d *Dog)Eat()  {
	fmt.Println("dog is eat..")
}

func main() {
	var d *Dog = &Dog{
		Feet:"feet feet feet..",
		Animal : &Animal{
			Name:"Dog",
			Sex:"男",
		},
	}
	d.Eat()
	d.Talk()
}




