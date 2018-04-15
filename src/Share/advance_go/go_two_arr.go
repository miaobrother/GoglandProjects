package main

import "fmt"

func main() {
	var two_arr [4][5]int
	k :=0
	for i := 0;i < 4; i ++{
		for j := 0;j <5;j++{
			k++
			two_arr[i][j] = k
		}
	}
	fmt.Printf("The two arr is %v\n",two_arr)


	bb := [4][5]int{{1,2,3,4,}, {6,7,8,9,10}, {11,12,13,14,15}, {16,17,18,19,20}}
	fmt.Printf("The init_two_arr is %v\n",bb)

}