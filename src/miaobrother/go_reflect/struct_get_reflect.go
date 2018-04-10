package main

import (
	"fmt"
	"reflect"
)

type S struct {
	A int
	B string
}

func (s *S) Test()  {
	fmt.Println("This is a test method.")
}

func (b *S) Best()  {
	fmt.Println("This is a Best method.")
}

func main() {
	s := S{21,"zhangsan"}
	v := reflect.ValueOf(&s)
	fmt.Println(v.Type())
	t := v.Type()
	fmt.Println("t is :",t)
	v.Elem().Field(0).SetInt(100)
	fmt.Println("method num :",v.NumMethod())
	fmt.Println("The change s after is:",s)
	for i := 0;i < v.NumMethod();i ++{
		f := t.Method(i)
		fmt.Printf("%d method, name:%v, type:%v\n",i,f.Name,f.Type)
	}
}
