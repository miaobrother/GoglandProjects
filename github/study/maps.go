package main

import "fmt"

func main() {
	// 定义一个变量 类型是map[int]string  type
	m := map[int]string{1: "python", 2: "golang", 3: "c++", 4: "lua"}

	for key, value := range m {
		fmt.Println(key, value)
	}
	value, ok := m[0]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("不存在")
	}
	delete(m, 1)
	fmt.Println(m)
	//map做函数参数

}
