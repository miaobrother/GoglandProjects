package main

import "fmt"


type People struct {
	Name string
	Age int
}

type Student struct {
	Score int
	People
}



func main()  {
	var s Student
	s.Name = "zhang"
	s.Age = 27
	s.Score = 100
	fmt.Printf("%#v\n",s)
}
