package main

import "fmt"

func testReverseStringOne() {
	var strOne = "Hello"
	var bytes []byte = []byte(strOne)
	fmt.Printf("This is a byte:%c", bytes) //bytes现在是一个切片

	for i := 0; i < len(bytes)/2; i++ {
		tmp := bytes[len(strOne)-i-1]
		bytes[len(strOne)-i-1] = bytes[i]
		bytes[i] = tmp
	}
	strOne = string(bytes)
	fmt.Println(string(strOne))
}

func main() {
	//testReverseStringOne()
	//testReverseStringTwo()
	testReversePalindrome()
}

func testReverseStringTwo() {
	var strTwo = "初庆峰是逗逼"
	var bytesTwo []rune = []rune(strTwo)
	fmt.Printf("This is a byte:%c\n", bytesTwo)

	for i := 0; i < len(bytesTwo)/2; i++ {
		tmp := bytesTwo[len(bytesTwo)-i-1]
		bytesTwo[len(bytesTwo)-i-1] = bytesTwo[i]
		bytesTwo[i] = tmp
	}

	strTwo = string(bytesTwo)
	fmt.Println(string(strTwo))
}

func testReversePalindrome() {

	var strThree = "noon"
	var bytesThree []rune = []rune(strThree)
	fmt.Printf("This is a byte:%c\n", bytesThree)

	for i := 0; i < len(bytesThree)/2; i++ {
		tmp := bytesThree[len(bytesThree)-i-1]
		bytesThree[len(bytesThree)-i-1] = bytesThree[i]
		bytesThree[i] = tmp
	}

	bytesFour := string(bytesThree)
	if bytesFour == strThree {
		fmt.Println(strThree, "is Palindrome")
	} else {
		fmt.Println(strThree, "not is Palindrome")
	}
	fmt.Println(string(bytesThree))

}
