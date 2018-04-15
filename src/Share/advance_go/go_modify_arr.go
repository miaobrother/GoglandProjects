package main

import "fmt"

//数组做函数的参数，是值传递
//实参数组的每个元素给形参数组的拷贝。并不影响原有的元组
func modify(abc [5]int)  {
	abc[0] = 911
	fmt.Printf("The modify after is %v\n",abc)//The modify after is [911 2 3 4 5]
}
func main() {
	abc := [5]int{1,2,3,4,5}
	modify(abc)
	fmt.Printf("The main is %v\n",abc)//The main is [1 2 3 4 5]
}