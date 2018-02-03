package main

import "fmt"

var a = "This is stored in the variable a"
var b ,c string  = "stored in b","stored in c"
var d string

func main() {
	d = "stored in d"

	var e = 42
	f := 43
	g := "stored in g"
	h,i := "stored in h","stored in i"
	j,k,l,m := 44.7,true,false,'m'
	n := "n"
	o := `o`
	var p = `p`

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)
	fmt.Printf("%T\n",o)
	fmt.Printf("%T\n",p)

	//var a int
	//var b string
	//var c float64
	//var d bool
	//
	//fmt.Printf(" a is:%v\n",a)
	//fmt.Printf("b is %v\n",b)
	//fmt.Printf(" c is:%v\n",c)
	//fmt.Printf("d is %v\n",d)
	//
	//fmt.Println()
	//var message = "Hello World!"
	// a , b ,c := 1,false,3
	// d := 4
	// e := true
	//fmt.Println(message,a,b,c,d,e)


}
