package main

import (
	"fmt"
	"os"
)
var (
	path string
)
func main() {
	fmt.Printf("Please enter your file:")
	fmt.Scanln(&path)
	{
		if _, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				fmt.Println("file does not exist:", path)
				fmt.Println()
				fmt.Println("您是否需要创建该目录:",path)

			}
		} else {
			fmt.Println("The file is Exist:", path)
		}

	}

}
