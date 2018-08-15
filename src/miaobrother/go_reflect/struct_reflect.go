package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
	Sex  string
}

func main() {

	var s Student
	v := reflect.ValueOf(s)
	t := v.Type()
	//fmt.Println(t.Kind())
	k := t.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("s is a int64\n")
	case reflect.Float64:
		fmt.Printf("s is a float64\n")
	case reflect.Struct:
		fmt.Printf("s is a struct\n")
		fmt.Printf("filed num of s is %d\n", t.NumField())
		fmt.Printf("filed num of s is %d\n", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			filed := v.Field(i)
			fmt.Printf("The s filed is : %v value is :%v\n", filed.Type(), filed.Interface())
		}
	}
}
