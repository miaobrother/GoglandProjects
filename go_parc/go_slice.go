package main

import "fmt"

func main() {

	s := make([]string,3)

	fmt.Println("emp:",s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:",s)
	fmt.Println("get:",s[2])
	fmt.Println("len:",len(s))

	c := make([]string,len(s))

	copy(c,s)

	fmt.Println("cpy:",c)

	l := s[2:]
	fmt.Println("l is:",l)


	t := []string{"go","h","i"}
	fmt.Println("dcl:",t)


	twoDD := make([][]int,3)
	for i := 0; i <3; i ++{
		innerLen := i +1
		twoDD[i] = make([]int,innerLen)
		for j := 0;j < innerLen; j ++{
			twoDD[i][j] = i +j
		}
	}
	fmt.Println("2d:",twoDD)
}
