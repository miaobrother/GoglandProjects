package main

import "fmt"

func main() {
	s1 := "abcd"
	b1 := []byte(s1)
	fmt.Println(b1)

	r1 := []rune(s1)
	fmt.Println(r1)


	s2 := "中文表示"

	b2 := []byte(s2)
	fmt.Println(b2)

	r2 := []rune(s2)
	fmt.Println(r2)
	fmt.Printf("%T\n",r2)
}