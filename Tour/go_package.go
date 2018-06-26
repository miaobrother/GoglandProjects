package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now().Unix()
	rand.Seed(t)
	fmt.Println("My favorite number is", rand.Int31())
}
