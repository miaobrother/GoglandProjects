package main

import (
	"fmt"
)

type User struct {
	Username string
	Age int
	Sex int
	AvatarUrl string
}

func main() {
	var user User
	user.Age = 18
	user.Sex = 1

	fmt.Printf("The user age is:%d\n",user.Age)
}

