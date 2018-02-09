package main

import (
	"fmt"
	//"crypto/tls"
)

type Vartex struct {
	x ,y int
}

var (
	v = Vartex{1,2}
	v1 = Vartex{x:1}
	v2 = Vartex{y:2}
	v3 = &Vartex{}

)

func main() {
	fmt.Println(v,v1,v2,*v3)
}

