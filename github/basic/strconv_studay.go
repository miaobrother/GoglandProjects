package main

import (
	"fmt"
	"strconv"
	//"os"
)

func main() {
	//i,err := strconv.Atoi("-42")
	//if err != nil{
	//	fmt.Fprintf(os.Stderr,"err:",err)
	//}
	//s:= strconv.Itoa(i)
	//fmt.Println(s)
	//b,err := strconv.ParseFloat("3.1415926",32)
	//if err != nil{
	//	fmt.Println("The error is:",err)
	//}else {
	//	fmt.Println(b)

	//b := []byte("bool:")
	//b = strconv.AppendBool(b,true)
	//fmt.Println(string(b))

	v := "10"
	if s,err := strconv.Atoi(v);err ==nil{
		fmt.Printf("%T,%v\n",s,s)
	}

}
