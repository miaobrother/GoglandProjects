package main

import "fmt"

func main() {
	info :=[4]string{
		"Zhang",
		"19",
		"python",
		"Go",
	}
	fmt.Println(info)
	a :=info[2:4]
	fmt.Println("a is :",a)


	a[0] = "Chen"
	fmt.Println(a)
	fmt.Println(info)
}