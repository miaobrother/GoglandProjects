package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}

func main() {
	fmt.Println(digitRegexp)
	fmt.Println(FindDigits("./digit"))

}
