package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	i,err := strconv.Atoi("-42")
	if err != nil{
		fmt.Fprintf(os.Stderr,"err:",err)
	}
	s:= strconv.Itoa(i)
	fmt.Println(s)

}