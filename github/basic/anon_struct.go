package main

import "fmt"

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human //匿名字段，那么默认的Student就包含了Human的字段
	speciality string
}

func main() {
	mark := Student{Human{"mark",25,130},"Computer"}
	fmt.Println("He name is:",mark.name)
	fmt.Println("He age is :",mark.age)
	fmt.Println("He weight is:",mark.weight)
	fmt.Println("He speciality is:",mark.speciality)

	mark.speciality = "AI"

	fmt.Println("He speciality changed after is:",mark.speciality)


	mark.weight = 140
	fmt.Println("He weight changed after is:",mark.weight)
}

