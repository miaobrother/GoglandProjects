package main

import "fmt"
//This is a reverse string pro

func reverse(str string) string  {
	rs := []rune(str)
	len := len(rs)
	var tt []rune

	tt = make([]rune,0)
	for i:= 0;i<len;i++ {
		tt = append(tt,rs[len -i -1])

	}
	//fmt.Println(tt)
	return string(tt[0:])
}

func main() {
	res := reverse("I Love Code")
	fmt.Println(res)
	//fmt.Println()
}
