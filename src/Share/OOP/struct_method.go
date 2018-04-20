package main

import (
	"fmt"
)

type User struct { //定义一个结构体
	Name string
	Email string
}

func (u *User) Notify()  {// 定义一个 Notify method
	fmt.Printf("%v : %v \n",u.Name,u.Email)
}

func main()  {
	uOne := User{"golang","zihao121908@gmail.com"}
	uOne.Notify()

	uTwo := User{"go","gogo.com"}
	uThree := &uTwo
	uThree.Notify()
}
