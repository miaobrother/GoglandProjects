package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	//"strconv"
)

func CreatNum(p *int) {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(1e4)
	if len(string(num)) > 4 {
		log.Fatal("The num is invailead")
	} else {
		fmt.Printf("the num is :%d\n", num)
	}
	*p = num
}

func main() {
	var randNum int
	CreatNum(&randNum)
	fmt.Printf("%d\n", randNum)
	a, b, c, d := randNum/1000, (randNum%1000)/100, (randNum%100)/10, randNum/1000
	fmt.Println(a, b, c, d)
}
