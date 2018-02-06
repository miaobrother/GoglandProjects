package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First string
	Last string
	Age int
}

func main()  {
	var p1 person
	//fmt.Println(p1.First)
	//fmt.Println(p1.Last)
	//fmt.Println(p1.Age)

	bs := []byte(`{"First":"James", "Last":"Bond","Age": 30}`)

	json.Unmarshal(bs,&p1)

	fmt.Println("------------------")

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	fmt.Printf("%T\n",p1)

}
