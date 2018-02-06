//func main() {
//	mySlice := []string{"Monday","Tuesday"}
//
//	myOtherSlice := []string{"Wednsday","Thursday","Friday"}
//
//	mySlice = append(mySlice,myOtherSlice...)
//
//	//fmt.Println("Slice is :",mySlice)
//
//	mySlice = append(mySlice[:2],mySlice[3:]...)
//	fmt.Println(mySlice)
//}

//func main() {
//	var student []string
//	//student := []string{}
//	var students [][]string
//
//	fmt.Println(student)
//	fmt.Println(students)
//	fmt.Println(student == nil)
//}

//package main
//
//import (
//"fmt"
//)
//
//func main() {
//
//	transactions := make([][]int, 0, 3)
//
//	for i := 0; i < 3; i++ {
//		transaction := make([]int, 0, 4)
//		for j := 0; j < 4; j++ {
//			transaction = append(transaction, j)
//		}
//		transactions = append(transactions, transaction)
//	}
//	fmt.Println(transactions)
//}

package main

import "fmt"

func main() {
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	fmt.Println(mySlice[0])
	mySlice[0]++
	fmt.Println(mySlice[0])
}