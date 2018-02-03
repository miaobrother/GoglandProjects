package main

import "fmt"

func Test() int  {
	return 100
}

func TestFor(){
	//var a int
	for a :=0; a < 10; a++ {
		fmt.Printf("Hehe is %d\n",a)
	}
}



func main()  {
	TestFor()
	//a :=Test()
	if a := Test();a >10{
		fmt.Printf("good %d\n",a)
	}else {
		fmt.Printf("The value is %d",a)
	}
}