package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input  = "i love you "
	in := bufio.NewScanner(strings.NewReader(input))
	in.Split(bufio.ScanWords)
	count := 0
	for in.Scan() {
		count++
	}


	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}
