package main

import "fmt"

type Humaner interface {
	//方法，只有声明，没有实现，只能由别的类型实现
	sayHi()

}

type Teacher struct {
	name string
	id int
}

func (t *Teacher) sayHi()  {
	fmt.Printf("Teacher [%s  %d] sayHi\n",t.name,t.id)
}


type Stu struct {
	score string
	age int
}

func (s *Stu) sayHi()  {
	fmt.Printf("Stu [%s  %d] sayHi\n",s.score,s.age)
}

type MyStr string

func (ttmp *MyStr) sayHi()  {
	fmt.Printf("MyStr [%s] sayHi\n",*ttmp)
}

func WhoSayHi(i Humaner)  {
	i.sayHi()
}
func main() {
	s := &Stu{"99",20}
	t := &Teacher{"lisi",20}
	var str MyStr = "hello world"

	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)
}


func mainOne() {
	var i Humaner
	//定义一个接口类型的变量

	//只是实现了此接口方法的类型，那么这个类型的变量就可以给i赋值
	s := &Stu{"99",19}
	i = s
	fmt.Println(i)

	t := &Teacher{"zhangsan",99}
	i = t
	i.sayHi()

}
