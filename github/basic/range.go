package main

import (
	"fmt"
)

//func main() {
//	a := []int{1,2,3,4}
//	for i,v := range a{
//		if i == 0{
//			a[1],a[2] = 999,999
//			fmt.Println(a)
//		}
//		a[i] = v + 100
//	}
//	fmt.Println(a)
//}
func main() {
	s := []int{1,2,3,4,5,6,7}
	for i,j := 0,len(s) -1; i < j;i,j = i +1,j-1{
		s[i],s[j] = s[j],s[i]
	}
	fmt.Println(s)
}