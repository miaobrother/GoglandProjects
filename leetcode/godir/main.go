package main

import (
	"fmt"
)

func testStr()  {
	var str = "hello我爱你中国"
	fmt.Printf("str[0]= %c\n",str[0])

	byteStr := []byte(str)
	byteStr[0] = '9'
	fmt.Printf("The byte is %v\n",string(byteStr))

	runeStr := []rune(str)
	for _,v := range runeStr{
		fmt.Println(string(v))
	}
	runeStr[6] ='很'
	fmt.Printf("The rune is %v\n",string(runeStr))
}
func main() {
	testStr()
}
