package main

import (
	"fmt"
)

type UserOne struct {
	Username string
	Age int
	Sex int
}

func main() {

	var user *UserOne
	fmt.Printf("user is :%v\n",user)// 不初始化的时候是  nil

	var user01 *UserOne = &UserOne{}
	user01.Age = 18
	user01.Sex = 1
	user01.Username = "zhangsan"

	fmt.Printf("The username is %#v\n",user01)//The username is &main.UserOne{Username:"zhangsan", Age:18, Sex:1}
}

