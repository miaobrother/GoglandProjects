package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("write\n",i,"as")

	switch i{
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	//当上面的case都没有满足的时候执行default所指定的逻辑块。
	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	t := time.Now()
	switch  {
	case t.Hour()<12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}