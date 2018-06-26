package main

import "fmt"

func Swap(x,y string) (string,string)  {
	return y,x

}

func main() {
	a ,b := Swap("Hello","Tina")
	fmt.Println(a,b)

}
