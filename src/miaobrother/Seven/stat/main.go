package main

import (
	"fmt"
	"bufio"
	"os"
	//"unicode"
)

func stat(str string) {
	var en_count int
	var sp_connt int
	var num_count int
	var other_count int
	utf8Arr := []rune(str)
	for _, v := range utf8Arr {
		if v >= 'a' && v < 'z' || v > 'A' && v <= 'Z' {
			en_count++
			continue
		}
		if v == ' ' {
			sp_connt++
			continue
		}
		if v >= '0' && v <= '9' {
			num_count++
		}
		other_count++
	}
	fmt.Printf("en_count=%d sp_count=%d num_count=%d other_count=%d", en_count, sp_connt, num_count, other_count)

}

func main() {
	//var str string

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("read string failed,err:%v\n", err)
		return
	}
	stat(line)

}
