package main

import (
	"fmt"
	"math/rand"
)

func testInterface() {
	var a interface{}
	var b int = 100
	var c float64 = 2.3
	var d string = "Hello"
	a = b
	fmt.Printf("The a is %v\n", a)
	a = c
	fmt.Printf("The a is %v\n", a)

	a = c
	fmt.Printf("The a is %v\n", a)
	a = d
	fmt.Printf("The a is %v\n", a)

}

func studentStore() {
	var stuMap map[int]map[string]interface{}
	stuMap = make(map[int]map[string]interface{}, 19)
	//插入学生id，姓名，分数，年龄
	var id = 1
	var name = "stu01"
	var score = 99.9
	var age = 19
	val, ok := stuMap[id]
	if !ok {
		val = make(map[string]interface{})
	}
	val["name"] = name
	val["id"] = id
	val["score"] = score
	val["age"] = age
	stuMap[id] = val

	fmt.Printf("stuMap:%#v\n", stuMap)
	for i := 0; i < 10; i++ {
		val, ok := stuMap[i]
		if !ok {
			val = make(map[string]interface{}, 16)
		}
		val["name"] = fmt.Sprintf("stu%d", i)
		val["id"] = id
		val["score"] = rand.Intn(100)
		val["age"] = rand.Intn(30)
		stuMap[id] = val
		fmt.Printf("stuMap:%#v\n", stuMap)
	}
}

func main() {

	//testInterface()
	studentStore()
}
