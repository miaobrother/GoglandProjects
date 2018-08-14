package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	en_count    int
	space_count int
	num_count   int
	other_count int
)

func statChar(str string) {
	utf8Arr := []rune(str)
	for _, v := range utf8Arr {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			en_count++
			continue
		}
		if v == ' ' {
			space_count++
			continue
		}
		if v >= '0' && v <= '9' {
			num_count++
			continue
		}
		other_count++
	}
	fmt.Println(en_count, space_count, num_count, other_count)
}

func main() {

	readerStrings := bufio.NewReader(os.Stdin)
	line, err := readerStrings.ReadString('\n')
	if err != nil {
		log.Fatal("Got a error:", err)
		return
	}
	statChar(line)
}
