package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	num := rand.Intn(100)
	fmt.Println("The num: ",num)

	for i:= 0;i < 100;i++{

		//fmt.Printf("The i is: %v \n",num)


		if i % 1 == 0 && i % 3 ==0 && i % 4 == 0 {
			continue
		}
		fmt.Printf("The i is: %v \n",i)
	}

}