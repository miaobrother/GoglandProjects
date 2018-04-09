package main

import (
	"fmt"
	"reflect"
)

type StudentOne struct {
	Name string
	Age int
	Sex string
}

func main() {

	var s StudentOne
	v := reflect.ValueOf(&s)
	v.Elem().Field(0).SetString("stu01")
	v.Elem().FieldByName("Sex").SetString("nan")
	v.Elem().FieldByName("Age").SetInt(99)
	fmt.Printf("s is %#v\n",s)

}
