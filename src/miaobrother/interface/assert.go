package main

import (
	"fmt"
	//"log"
	//"log"
)

func test(a interface{})  {//定义的一个空接口
	//s ,ok := a.(int)
	//if !ok{
	//	log.Fatal(ok)
	//	return
	//}
	//fmt.Println(s)
	strOne,ok := a.(int)
	if ok{
		fmt.Println(strOne)
		return
	}


}

func testSwitch(a interface{})  {
	switch v := a.(type){
	case string:
		fmt.Printf("a is string, value:%v\n",v)
	case int:
		fmt.Printf("a is int, value:%v\n",v)
	case float32:
		fmt.Printf("a is float32, value:%v\n",v)
	case float64:
		fmt.Printf("a is float64, value:%v\n",v)
	default:
		fmt.Printf("No suport type is %v\n",v)
	}
}

func main() {

	//var a int
	//test(a)


	var b string = "Hello"
	//test(b)
	testSwitch(b)
}