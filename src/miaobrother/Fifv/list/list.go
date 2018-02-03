package main

import "fmt"

type Teacher struct {
	Name string
	Age int
	Next *Teacher
}


func createList(){
	var header *Teacher = &Teacher{}
	header.Age = 200
	header.Name = "a"

	printlist(header)
}

func printlist(h *Teacher) {
	for h != nil{
		fmt.Printf("Name:%v Age:%v\n",h.Name,h.Age)
		h =h.Next
	}
}

