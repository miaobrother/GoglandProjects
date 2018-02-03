package main

import "fmt"

func main() {

	a := 10
	b := "Golang"
	c := 3.14
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	//%T  是查看是什么类型。
	fmt.Printf(" a is :%T \n",a)
	fmt.Printf(" b is :%T \n",b)
	fmt.Printf(" c is :%T \n",c)
	fmt.Printf( "d is :%T \n",d)
	fmt.Printf("e is :%T \n",e)
	fmt.Printf("f is :%T \n",f)
	fmt.Printf("g is :%T \n",g)
}
