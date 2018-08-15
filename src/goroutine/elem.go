package main

import (
	"fmt"
	"reflect"
)

func main() {

	var b int
	b = 200
	testInt(&b)
	fmt.Println(b)
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	fmt.Println("the val is ", &val)
	val.Elem().SetInt(1000)
	c := val.Elem().Int()
	fmt.Printf("get value interface{} %d\n", c)
	fmt.Printf("string val:%d\n", val.Elem().Int())
}
