package main

import "fmt"

const  a   =  1
const b = "str_const"

const(
	c,d,e = 1,2,3
	f,g,h = 4,5,6
)

const i,j,k = 99,88,33

const t = iota

const(
	l = iota
	m
	n
)

func main() {
	fmt.Println(a,b)
	fmt.Println(c,d,e,f,g,h)
	fmt.Println(i,j,k)
	fmt.Println(t,l,m,n)
}


