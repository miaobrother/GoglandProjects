package basic

import (
	"fmt"
)

func updateSlice(s []int) {
	s[0] = 100
}
func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[2:6]
	s2 := arr[:7]
	fmt.Println(s)
	fmt.Println(s2)

	updateSlice(s)
	fmt.Println(arr)
	fmt.Println(s)
	fmt.Println(s2)

	fmt.Println("Extending slice")

	s3 := arr[2:6]
	fmt.Println(s3)
	s4 := s3[3:5]
	fmt.Println(s4)
}
