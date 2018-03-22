package main

import (
	"fmt"
	"math/rand"
	"encoding/json"
)

type User7 struct {
	UsernameTwo string `json:"username"`
	SexTwo string `json:"sex_two"`
	Score float32
}

func main()  {
	user := &User7{
		UsernameTwo:"user01",
		SexTwo:"男",
		Score:rand.Float32()* 100,
	}

	data,_:= json.Marshal(user)
	fmt.Printf("json str :%s\n",string(data))
}

