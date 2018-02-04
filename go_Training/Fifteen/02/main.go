//This is a Test
package main

import "fmt"

func max(numbers ...int) int  {
	var largest int
	for i,v := range numbers{
		if v > largest || i == 0{
			largest = v
		}
	}
	return largest
}

func main() {
	//greatest := max(1000,-1,9,13,145,65,467,989)
	greatest := max(-200,-700)
	fmt.Println(greatest)
}