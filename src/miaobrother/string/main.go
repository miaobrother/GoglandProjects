package main

import (
	"fmt"
	//""
)

func testString() {
	var str = "Hello"
	fmt.Printf("str[0] = %c\n str len is %d\n", str[0], len(str))
	for i, v := range str {
		fmt.Println(i, v)
		fmt.Printf("i is %d,v is %c", i, v)
	}
}

func calc(s string) (charCount int, numCount int, spaceCount int, otherCount int) {
	utfChars := []rune(s)

	for i := 0; i < len(utfChars); i++ {
		if utfChars[i] >= 'a' && utfChars[i] <= 'z' || utfChars[i] >= 'A' && utfChars[i] <= 'Z' {
			charCount++
			continue
		}
		if utfChars[i] >= '0' && utfChars[i] <= '9' {
			numCount++
		}
		if utfChars[i] == ' ' {
			spaceCount++
			continue
		}
		otherCount++
	}
	return
}

func calcCharCount() {
	var s string = "ni hao wo 不知道什么,.kaka0a9191  "
	charCount, numCount, spaceCount, otherCount := calc(s)
	fmt.Printf("The CharCount is :%d\n", charCount)
	fmt.Printf("The CharCount is :%d\n", numCount)
	fmt.Printf("The spaceCount is %d\n" ,spaceCount)
	fmt.Printf("The otherCount is %d\n", otherCount)

	//calc(s)
}

func main() {
	//testString()
	//calc("ni hao wo 不知道什么,.kaka0a9191  ")
	calcCharCount()
}
