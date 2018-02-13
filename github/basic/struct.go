package main

import (
	"fmt"
	//"crypto/x509/pkix"
)

type person struct {
	name string
	age int
}

func Older(p1,p2 person)(person,int)  {
	if p1.age > p2.age{
		return p1,(p1.age-p2.age)
	}
	return p2,p2.age-p1.age
}

func main() {
	var tom person

	tom.name,tom.age = "Tom",19

	bob := person{age:20,name:"Bob"}

	//paul := person{"zhangsan",21}

	tb_Older,tb_diff := Older(tom,bob)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)
}
