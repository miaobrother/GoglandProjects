package main

import "fmt"

func main() {

	fmt.Println(green("Jack","Zhangsan"))
}

func green(fname ,lname string) (s string) {

	s = fmt.Sprint(fname,lname)
	return
	//return fmt.Printf("first is %s\n %s\n",fname,lname)
}