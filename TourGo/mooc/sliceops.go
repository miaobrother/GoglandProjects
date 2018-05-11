package main

import (
	"fmt"
)

func calcSlice(s []int)  {
	fmt.Printf("The s len = %d,and cap is cap = %d\n",len(s),cap(s))
}
func main() {

	var s []int
	for i := 0; i < 100; i++ {
		calcSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Printf("The slice s is %v\n", s)

	s1 := make([]int,10,20)
	s3 := append(s1, 1,2,3)
	s4 := append(s3[:10] ,s3[12:]...)
	fmt.Println("The s4 is ",s4)
	fmt.Println(s3)
	s2 := make([]int,10,32)
	if cap(s1) == cap(s2){
		fmt.Println("Eq")
	}
	fmt.Printf("The s1 is %#v\n",s1)
	fmt.Printf("The s2 is %#v\n",s2)

}
