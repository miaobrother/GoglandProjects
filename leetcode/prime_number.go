package main

import "fmt"

func Justify(n int) bool {
	if n <= 1 {
		return false
	}
	for k := 2; k < n; k++ {
		if n % k == 0 {
			return false
		}

	}

	return true
}

func IsPrimeNumber() {
	for i := 2; i < 100; i++ {
		if Justify(i) == false {
			fmt.Printf("%d is a prime\n", i)
		}

	}

}

func main() {
	IsPrimeNumber()

}
