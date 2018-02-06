package main

import "fmt"

func main() {
	buckets := make([]int,1)

	buckets[0] = 42

	fmt.Println(buckets[0])
	buckets[0]++
	fmt.Println(buckets[0])
	fmt.Printf("%T\n",buckets)
}