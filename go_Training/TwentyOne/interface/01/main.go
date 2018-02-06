package main

import "fmt"


type square struct {//define a struct
	side float64
}

func (z square) area() float64  {
	return z.side * z.side
}

func main(){
	s := square{20}

	fmt.Println("Area:",s.area())
}
