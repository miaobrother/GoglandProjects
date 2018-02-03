package main

import "fmt"

func main() {
	names := [4]string{
		"Jone", //0
		"Jack", //1

		"Alex", //2
		"Erin", //3
	}
	a := names[0:2]
	b := names[1:3]
	//fmt.Println(a,b)
	fmt.Printf("The A is:%v\n",a)
	fmt.Printf("The B is:%v\n",b)

	b[0] = "XXX"
	fmt.Println(a,b)
	//fmt.Println(names)
}
