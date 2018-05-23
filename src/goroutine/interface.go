package main

import "fmt"

type UserThree struct {
	id int
	name string
}

type ManagerOne struct {
	UserThree
}

func (self *UserThree) ToString() string  {
	return fmt.Sprintf("UserThree: %p, %v",self,self)
}

func main() {
	m := ManagerOne{UserThree{1,"lisi"}}

	fmt.Printf("Manager :%p\n",&m)

	fmt.Println(m.ToString())
}
