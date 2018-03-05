package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//使用scanner来包裹无缓冲的os.Stdin

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil{
		fmt.Fprintln(os.Stderr,"error:",err)
		os.Exit(1)
	}
}