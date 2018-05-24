package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	var t = T{
		10,
		"zhangsan",
	}

	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	fmt.Printf("The typeOfT is %v\n",typeOfT)
	fmt.Printf("The s is %v\n",s)

	for i := 0; i < s.NumField(); i++{
		f := s.Field(i)
		fmt.Printf("%d:%s %s = %v\n",i,typeOfT.Field(i).Name,f.Type(),f.Interface())
	}


	s.Field(0).SetInt(99)
	s.Field(1).SetString("sunset strip")
	fmt.Println("t is now:",t)

}
