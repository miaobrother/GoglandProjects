package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sliceToMap() {
	rand.Seed(time.Now().UnixNano())
	var s []map[string]int
	s = make([]map[string]int, 5, 19)

	for k, v := range s {
		fmt.Printf("The key is %d, The value is %v\n", k, v)
	}
	s[0] = make(map[string]int, 5)
	s[0]["student01"] = 1000
	s[0]["student02"] = 2000
	s[0]["student03"] = 3000
	for k, v := range s {
		fmt.Printf("The key is %d, The value is %v\n", k, v)
	}
}

func mapToSlice() {
	rand.Seed(time.Now().UnixNano())
	var m map[string][]int     //定义一个map，value 是slice的
	m = make(map[string][]int) //初始化
	key := "stu01"             //先定义一个key
	value, ok := m[key]        //判断key是否存在
	if !ok {
		m[key] = make([]int, 0, 18) //如果不存在的话则初始化一下slice
		value = m[key]
	}
	value = append(value, 100)
	value = append(value, 200)
	value = append(value, 300)
	m[key] = value

	fmt.Printf("The mapToSlice is %v\n", m)
	/*
		for k,v := range m{

			fmt.Printf("The key is %v The va is %v\n",k,v)
		}
	*/
}
func main() {
	//sliceToMap()
	mapToSlice()

}
