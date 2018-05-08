package main

import "fmt"

type UserOne struct {
	id int
	name string
}

type Manager struct {
	UserOne
}

func (self *UserOne) ToString() string  {
	return fmt.Sprintf("User:%p,%v",self,self)
}

func main() {
	m := Manager{UserOne{1,"zhangsan"}}

	fmt.Printf("manager: %p\n",&m)
	fmt.Println(m.ToString())
}