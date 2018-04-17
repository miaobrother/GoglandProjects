package main

import "fmt"

type PersonF struct {
	name string
	age int
}

func (tmp PersonF) PrintInfo()  {
	fmt.Println("tmp = ",tmp)
}

//通过一个函数，给成员变量赋值

func (p *PersonF) SetInfo(n string,a int)  {
	p.name = n
	p.age = a
}
func main() {
	//定义一个初始化参数
	p := PersonF{"Chen",19}
	p.PrintInfo()

	//定义一个结构体变量
	var pTwo PersonF
	(&pTwo).SetInfo("Jack_ZHang",26)
	pTwo.PrintInfo()

}
