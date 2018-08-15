package main

import "fmt"

type Teacher struct {
	Name string
	Age  int
	Next *Teacher
}

//func crite(){
//	var header = &Teacher{}
//	//fmt.Printf("the teacher is %#v\n",header)
//
//	header.Name = "n"
//	header.Age = 18
//
//	p := new(Teacher)
//	p.Name="abc"
//	p.Age = 17
//	header.Next = p
//	fmt.Printf("The p is %p\n",p)
//	printlist(header)
//}

func createinitheader(h *Teacher, name string, age int) *Teacher {
	p := &Teacher{}
	p.Name = name
	p.Age = age
	p.Next = h
	return p
}
func printlist(t *Teacher) {
	for t != nil {
		fmt.Printf("The t is %#v\n", t)
		t = t.Next
	}
}
func testcreateinit() {
	var header *Teacher
	//fmt.Printf("The hearder is %#v\n",header)
	header = createinitheader(header, "zhang", 19)
	header = createinitheader(header, "chen", 18)
	header = createinitheader(header, "li", 17)
	header = createinitheader(header, "wang", 16)
	printlist(header)
}

func main() {
	//crite()
	testcreateinit()

}
