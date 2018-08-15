package main

import (
	"fmt"
	"regexp"
)

const text = "My email is zihao@gmail.com"

func main() {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@.+\..+`)
	match := re.FindString(text)
	fmt.Println(match)

}
