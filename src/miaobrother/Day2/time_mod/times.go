package main

import ("fmt"
		"time"
)

const(
	Man = 1
	Famale = 2
)

func main()  {

	now := time.Now()
	second := now.Unix()
	fmt.Println("Now is ",now)
	if second % Famale == 0{
		fmt.Println("Fem",second)
	}else{
		fmt.Println("Man",second)
	}


}