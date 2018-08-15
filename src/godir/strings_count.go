package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var countMaps map[string]int

func statChar(str string) {
	countMaps = make(map[string]int)
	utf8Arr := []rune(str)
	for _, v := range utf8Arr {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			countMaps["en_count"]++
			continue
		}
		if v == ' ' {
			countMaps["space_count"]++
			continue
		}
		if v >= '0' && v <= '9' {
			countMaps["num_count"]++
			continue
		}
		countMaps["other_count"]++
	}
	for k, v := range countMaps {
		fmt.Printf("The key is:%v The value is:%v\n", k, v)
	}
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
