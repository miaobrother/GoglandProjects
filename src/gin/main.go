package main

import (
	"sort"
	"fmt"
	"math/rand"
	"time"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

type StudentSlice []*Student

func (s StudentSlice) Len() int {
	return len(s)
}

func (s StudentSlice) Less(i,j int) bool   {
	return s[i].Age > s[j].Age
}


func (s StudentSlice) Swap(i,j int)   {
	s[i].Age,s[j].Age = s[j].Age,s[i].Age
}
func main() {
	var studentSlice StudentSlice
	t := time.Now().UnixNano()
	rand.Seed(t)

	for i := 0;i < 10;i ++{
		var ss = &Student{
			Name:fmt.Sprintf("张 %d\n",i),
			Age:rand.Intn(30),
			Score:rand.Float32() * 100,
		}
		studentSlice = append(studentSlice,ss)
	}
	sort.Sort(studentSlice)
	for i := 0;i < len(studentSlice);i ++{
		fmt.Printf("%#v\n",studentSlice[i])
	}
}