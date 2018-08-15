package main

import (
	"fmt"
)

func main() {
	abc := []int{1, 1, 2}
	for i := 0; i < len(RemoveMap(abc)); i++ {
		fmt.Println(RemoveMap(abc)[i])
	}
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
	return tmp

}
