package main

import (
	"fmt"
)

func main() {
	abc := []int{1, 1, 2, 3, 4, 5, 6, 6}
	//for i := 0; i < len(RemoveMap(abc)); i++ {
	//}
	fmt.Println(RemoveMap(abc))
}

func RemoveMap(abc []int) (s []int) {
	var tmp []int
	for i, val := range abc {
		if i+1 < len(abc) {
			if abc[i] != abc[i+1] {
				tmp = append(tmp, val)

			}

		} else {
			tmp = append(tmp, val)
		}

	}
	//fmt.Println(tmp)
	return tmp

}
