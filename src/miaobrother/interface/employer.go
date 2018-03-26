package main

import "fmt"

func NewProgramer(name string,base float32,extend float32)Programer  {
	return Programer{
		name:name,
		base: base,
		extend :extend,

	}
}

func NewSale(name string,base float32,extend float32)Sale  {
	return Sale{
		name:name,
		base: base,
		extend :extend,

	}
}
type Employer interface {
	CalcSalary() float32
}

type Programer struct {//程序员
	name  string
	base float32
	extend float32
}

func (p Programer) CalcSalary() float32  {
	return p.base
}

type Sale struct {//销售
	name string
	base float32
	extend float32
}

func (p Sale) CalcSalary() float32  {
	return p.base + p.extend*p.base*0.5
}

func calcAll(e []Employer) float32  {
	var cost float32
	for _,v := range e{
		cost = cost + v.CalcSalary()
	}
	return cost
}

func main(){
	p1 := NewProgramer("Lisi",15000,0.04)
	//p2 := NewProgramer("Zhangs",14000,0.03)
	//p3 := NewProgramer("haha",1500,0.05)
	//
	//s1 := NewSale("wo ha",7000,0.9)
	//s3 := NewSale("wo ha",8000,0.9)

	var epl []Employer
	epl = append(epl,p1)
	cost := calcAll(epl)
	fmt.Printf("This mouth cost is :%f\n",cost)
	//fmt.Println(epl)


}
