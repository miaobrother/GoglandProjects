package main

import (
	"os"
	"fmt"
)

func testWriteFile() {
	file, err := os.OpenFile("/Users/zhangyalei/Desktop/123.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	for i := 0; i < 10; i++ {
		file.WriteString(fmt.Sprintf("Hello Wrold %d\n", i))
	}
}

func main() {
	testWriteFile()
}
