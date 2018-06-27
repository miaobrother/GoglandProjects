package main

import (
	"fmt"
	"time"
)
func main() {
	fmt.Println("When's Saturday?")

	today := time.Now().Weekday()
	switch time.Saturday {
	case today +0:
		fmt.Println("today is saturday")
	case today +1 :
		fmt.Println("today is unknow")
	case today +2 :
		fmt.Println("today is buzhidao")
	case today +3 :
		fmt.Println("not yet saturday")
	}
	fmt.Println(today)
}
