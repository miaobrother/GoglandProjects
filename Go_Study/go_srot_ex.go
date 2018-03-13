package main

import (
	"sort"
	"math/rand"
	"fmt"
)

type Student struct {
	name string
	age int
	score float64
}

type StudentSlice []*Student

func (p StudentSlice) Len()int  {
	return len(p)
}

func (p StudentSlice) Less(i,j int) bool  {
	return p[i].score > p[j].score
}

func (p StudentSlice) Swap(i,j int){
	p[i],p[j]  =  p[j],p[i]
}

func main() {
	var studentArr StudentSlice
	for i:=0;i <10;i ++{
		var s = &Student{
			name : fmt.Sprintf("张%d",i),
			age : rand.Intn(100),
			score: 100 * rand.Float64(),
		}
		studentArr = append(studentArr,s)
	}
	sort.Sort(studentArr)
	for i:=0;i <len(studentArr);i ++{
		fmt.Printf("%#v\n",studentArr[i])
	}

}

