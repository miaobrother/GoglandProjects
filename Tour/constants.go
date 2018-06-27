package main

import (
	"fmt"
	"math/rand"
	"time"
)
const Pi = 3.15

func main() {
	t := time.Now().Unix()
	rand.Seed(t)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Hello", Pi, "Days")
}
