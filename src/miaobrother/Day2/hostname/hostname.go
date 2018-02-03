package main

import ( "fmt"
		 "os"
)

func swap (a int,b int) (int ,int) {
	var c int
	c = a
	a = b
	b = c
	return a,b
}

func main() {
	name, ret := os.Hostname()
	fmt.Printf("%s\n %v\n", name, ret)
	val := os.Getenv("PATH")
	fmt.Printf("%s\n", val)

	a,b := 100,200
	a,b = swap(a,b)
	fmt.Printf("a = %d\n b = %d\n",a,b)
}

