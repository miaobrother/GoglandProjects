package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	eular()
	triangle()
	enums()
}

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func triangle() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Printf("the c is %v\n", c)
}

func eular() {
	fmt.Printf("%.3f,", cmplx.Exp(1i*math.Pi)+1)

	c := 3 + 4i

	fmt.Println(cmplx.Abs(c))

}

func enums() {
	const (
		python = iota
		java
		c
		swift
	)
	fmt.Println(python, java, c, swift)

}
