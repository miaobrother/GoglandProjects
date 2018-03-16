package main

import (
	"fmt"
)

func justIfy(n int) bool  {
	if n <= 1{
		return false
	}
	for i:=2;i < n;i++{
		if n % i == 0{
			return false
		}

	}
	return true

}
func isOrNotPrime()  {
	for i:= 2;i < 100;i ++{
		if justIfy(i) == true{
			fmt.Printf("%d is prime\n",i)
		}

	}
}
func main() {
	isOrNotPrime()
	
}
