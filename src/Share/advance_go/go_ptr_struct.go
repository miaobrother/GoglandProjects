package main

import "fmt"

type Ptr struct {
	id int
	name string
	sex byte
}

func main() {
	//指针有合法指向后，才能操作成员
	//先定义一个普通结构体变量
	var s Ptr


	var p *Ptr

	p = &s
	p.id = 2
	//s.id=3
	p.name= "zhangsan"
	//s.name = "chenmiao"
	p.sex = 'm'
	//s.sex ='f'
	fmt.Printf("The p is %#v\n",p)
	fmt.Printf("The s is %#v\n",s)


	//通过 new申请一个结构体

	p2 := new(Ptr)
	p2.sex ='f'
	p2.id =3
	p2.name ="limeimei"
	fmt.Println("The p2 is ",*p2)

}
