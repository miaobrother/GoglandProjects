package main

import (
	"fmt"
)

func main() {

	var mapOne map[string]string //声明Map后不能直接使用， 需要make一下；

	mapOne = make(map[string]string)
	mapOne["name"] = "Zhangsan"
	mapOne["Age"] = "25"

	mapTwo := make(map[string]string) //直接声明时使用make创建；
	mapTwo["Work"] = "Computer"

	//初始化 + 声明 + 赋值一体化
	mapThree := map[string]string{
		"e": "eee",
		"f": "ffff",
	}
	fmt.Println(mapOne)
	fmt.Println(mapTwo)
	fmt.Println(mapThree)

	v_map := map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
	}
	fmt.Println(len(v_map))

	for k, v := range v_map {
		fmt.Println(k, v)
	}

}
