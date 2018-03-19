package main

import (
	"fmt"
)

func insert_Sort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i;j >0;j--{
			if a[j] < a[j-1]{
				a[j],a[j-1] = a[j-1],a[j]
			}else {
				break
			}
		}


	}
	return a

}
func main() {
	var i []int = []int{2, 1, 9, 4, 8, 3, 0, 6}
	//fmt.Println(i)
	j := insert_Sort(i)
	fmt.Println(i)
	fmt.Println(j)
}
