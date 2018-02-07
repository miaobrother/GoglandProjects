package main

import (

	"fmt"
	//"sort"
	"sort"
)

type people []string

func (p people) Len() int  {
	return len(p)
}

func (p people) Swap(i,j int)  {
	p[i],p[j] = p[j],p[i]
}

func (p people) Less(i,j,int int) bool  {
	return p[i] < p[j]
}

func main() {
	studyGroup := people{"Zhang","Li","Chen","Yang"}

	fmt.Println(studyGroup)
	sort.Sort(sort.StringSlice(studyGroup))//报错。
	fmt.Println(studyGroup)
}