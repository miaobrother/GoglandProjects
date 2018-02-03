package main

import "fmt"

type User struct {
	name string
	flag uint8
/*	IsHongMing bool
	IsDaRen bool
	IsVip bool
	IsSuperVip bool
	IsBlue bool
	IsRed bool
*/
}
/*
func settings(user User,flag  bool)  {
	user.IsRed = flag

}
func IsRed(user User) bool  {
	return user.IsRed
}
*/

func set_HongMing(user User,isSet bool) User  {
	if isSet == true {
		user.flag = user.flag | 1
	}else {
		user.flag = user.flag | 0
	}
	return user
}
func is_HongMing(user User) bool  {
	result :=user.flag & 1
	return result == 1
}
func weibo_test(){
	var user User
	user.name = "test01"
	user.flag = 0
	result :=is_HongMing(user)
	fmt.Printf("user is HongMing:%t\n",result)

	user = set_HongMing(user,true)
	 result = is_HongMing(user)
	fmt.Printf("user is HongMing:%t\n",result)
}