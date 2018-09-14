package main

import (
	//"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	//"encoding/xml"
)

func main() {
	//string to  int

	str := "30kk"
	intValue ,_ := strconv.Atoi(str)
	fmt.Println(reflect.TypeOf(intValue))//利用反射看intValue的类型
	fmt.Println("The valueof:",reflect.ValueOf(intValue))

	//string to int64

	int64Value ,_ := strconv.ParseInt(str,10,64)
	fmt.Println(reflect.TypeOf(int64Value))



	//int to string

	intTmp := 100
	strTmp := strconv.Itoa(intTmp)
	fmt.Printf("The strTmp type is :%T\n",strTmp)



	//int64 to string

	 var intTmp64 int64
	 intTmp64 = 0xA
	 strTmp = strconv.FormatInt(intTmp64,10)
	 fmt.Println(reflect.TypeOf(strTmp))
}