package main

import (
	"fmt"
	"runtime"
)
func main() {
	//fmt.Println("Go runs on :")
	fmt.Println(runtime.GOOS)
	switch os := runtime.GOOS;os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("%s.",os)
	}
}