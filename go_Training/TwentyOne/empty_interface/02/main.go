package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func specs(a interface{})  {

	fmt.Printf("%T\n",a)
	fmt.Println(a)
}

func main() {
	fido := dog{animal{"woof"},true}

	fidi := cat{animal{"meow"},true}

	specs(fidi)
	specs(fido)
}