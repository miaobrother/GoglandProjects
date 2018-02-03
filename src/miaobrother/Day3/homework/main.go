package main

import "fmt"

func main()  {
	//for i := 101;i <= 200; i++ {
	//	if isPrime(i){
	//		fmt.Println(i)
	//	}
	//
	//}
	for i:=100;i<999;i++ {
		isNarc(i)
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2;i <n; i ++ {
		if n % 2 == 0{
			return false
		}
	}
	return  true
}

func isNarc(n int)bool  {
	i,j,k := n /100,n/10 %10, n %10
	fmt.Println("",i,j,k)
	return false
	}