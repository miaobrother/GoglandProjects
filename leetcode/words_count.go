package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input string

func main() {
	input := bufio.NewReader(os.Stdin)
	data, _, _ := input.ReadLine()

	out := strings.Fields(string(data))

	for i := 0; i < len(out); i++ {
		fmt.Println(out[i])
	}

}
