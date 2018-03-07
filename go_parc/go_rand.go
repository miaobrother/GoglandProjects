package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//EX rand.Ittn 返回一个整型随机数n, 0<=n<100
	fmt.Printf("The int is:%d\n", rand.Intn(100))
	fmt.Println()
	//0.0 <= f <1.0
	fmt.Printf("This float is:%v\n", rand.Float64())

	fmt.Print((rand.Float64() * 5) + 5,",")
	fmt.Print((rand.Float64() * 5) + 5)

	fmt.Println()

	s1 := rand.NewSource(42)
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100),",")
	fmt.Print(r1.Intn(100))

	fmt.Println()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100),",")
	fmt.Print(r2.Intn(100))
}
