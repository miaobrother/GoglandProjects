package basic

import (
	"fmt"
	//"reflect"
)

func printArray(arr [5]int) {

}

func main() {

	var arr [5]int
	fmt.Println(arr)

	var arrOne = [...]int{1, 2, 3, 4}
	for i := range arrOne {
		fmt.Println(arrOne[i])
	}

}
