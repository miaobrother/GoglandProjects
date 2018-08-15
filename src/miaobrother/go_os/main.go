package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		var buf[16]byte
		//os.Stdin.Read(buf[:])
		fmt.Println(string(buf[:]))
		os.Stdout.WriteString(string(buf[:]))
	*/
	var a int
	var b string
	var c float32

	//fmt.Fscanf(os.Stdin,"%d%s%f",&a,&b,&c)
	fmt.Println(a, b, c)

}
