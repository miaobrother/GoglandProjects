package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	fmt.Println("The t is n ", t)
	rand.Seed(t)
	fmt.Println("My favorite number is", rand.Intn(10))
}
