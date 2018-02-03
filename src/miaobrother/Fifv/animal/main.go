package main

import "fmt"

type Animal interface {
	Eat()
	Talk()
}

type Dog struct {

}

func (d *Dog) Eat()  {
	fmt.Println("Dog eating")
}
func (d *Dog) Talk()  {
	fmt.Println("Dog Talk.")
}

func main()  {
	var a Animal
	//a.Eat()
	var d Dog
	d.Eat()
	a = &d
	a.Eat()
	a.Talk()
}
