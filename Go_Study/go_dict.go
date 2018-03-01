package main

import "fmt"

func main() {
	m := make(map[string]int)//定义及初始化一个字典

	m["k1"] =7
	m["k2"] = 9
	fmt.Println("map:",m)

	v1 := m["k1"] //get value
	fmt.Println(v1)

	fmt.Println(len(m))

	delete(m,"k2")
	fmt.Println(m)

	_,ok := m["k2"]
	fmt.Println("ok:",ok)


	n := map[string]int{"foo":1,"bar":2}

	 fmt.Println("map:",n)
}
