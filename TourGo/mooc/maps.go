package main

import "fmt"

func main() {
	m := map[string]string{
		"name":"zhangsan",
		"age": "32",
		"school":"HanDan",
		"hobby":"a",
	}
	fmt.Printf("The m is %#v\n",m)

	m2 := make(map[string]string,10)
	m3 := make(map[string]string,10)
	if len(m2) == len(m3){
		fmt.Println( "M2 = m3")
	}
	fmt.Printf("The m2 is %#v\n",m2)
	fmt.Printf("The m2 is %#v\n",m3)

	for k,v := range m{
		fmt.Printf("The key is :%s, value is :%s\n",k,v)
	}

	custname ,ok := m["name"]
	if ok != true{
		fmt.Println("The ame not in m")
	}
	fmt.Println(len(custname))
}