package main

import "fmt"

func TwoSum(b [5]int,target int)  {
	for i := 0;i <len(b);i ++{
		other := target - b[i]
		for j := i+1;j < len(b); j++{
			if b[j] == other{
				fmt.Printf("%d,%d\n",i,j)
			}
		}
	}
}

func testTwoSum()  {
	abc := [...]int{1,3,5,8,7}
	TwoSum(abc,8)

}
func main() {
	testTwoSum()
}
