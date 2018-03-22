package main

import (
	"fmt"
)

type People struct {
	Name string
	Country string
}

func (p People)Print()  {//p 是People的实例
	fmt.Printf("name = %s country s %s\n",p.Name,p.Country)


}

func (p People)Set(name string,country string)  {
	p.Name = name
	p.Country = country
}

func main() {
	var p1 People = People{
		Name:    "people01",
		Country: "China",
	}
	p1.Print()
	p1.Set("people02","japen")
	p1.Print()
}

