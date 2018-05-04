// TestFlag project main.go
package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("name","","Input your username")

	flag.Parse()
	fmt.Println("hello,",*username)




}
