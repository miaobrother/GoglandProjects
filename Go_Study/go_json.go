package main

import (
	"encoding/json"
	"fmt"
	//"os"
)

type responseOne struct {
	Page int
	Fruits []string
}

type responseTwo struct {
	Page int       `json:"page"`
	Fruits []string  `json:"fruits"`
}

func main() {

	bolB,_ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))


	fltB,_ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB,_ := json.Marshal("gopher")
	fmt.Println(string(strB))




	sliD := []string{"apple","peach","pear"}

	sliB,_ := json.Marshal(sliD)
	fmt.Println((string(sliB)))





	mapD := map[string]int{"apple":5,"leetcode":7}

	mapB ,_ := json.Marshal(mapD)

	fmt.Println(string(mapB))


	res1D := &responseOne{
		Page:1,
		Fruits:[]string{"apple","peach","pear"},
	}

	res1B,_ := json.Marshal(res1D)

	fmt.Println(string(res1B))





	byt := []byte(`{"num":6.15,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt,&dat); err!= nil{
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"]
	fmt.Printf("%T\n",num)
}

