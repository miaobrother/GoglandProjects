package main

import (
	"fmt"
	//"testing"
)

func testInput()  {

	var a int
	var s string
	var f float32
	fmt.Scanf("%d%s%f\n",&a,&s,&f)

	fmt.Printf("The a is:%d,b is:%s,f is:%f\n",a,s,f)

}

func testScan()  {
	var a int
	var s string
	var f float32
	fmt.Scan(&a,&s,&f)
	fmt.Printf("The a is:%d,b is:%s,f is:%f\n",a,s,f)
}

func main()  {
	//testInput()
	testScan()
}
