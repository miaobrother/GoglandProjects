package main

import (
	"fmt"
	"strings"

)

func StrOperation()  {
	str := "hello"
	fmt.Printf("len:%d\n",len(str))
	str2 := str + " " +"world"
	fmt.Printf("str is %s\n",str2)
	str3 :="a,b,c,d,e,f,ghijklmnopqreto"
	result := strings.Split(str3,",")
	fmt.Printf("ret is %s\n",result)

	str5 :=strings.Join(result,".")
	fmt.Printf("str5 %s\n",str5)

	isContain := strings.Contains(str3,"o")
	fmt.Printf("contain:%t\n",isContain)

	str4 := "baidu.com"
	index :=strings.Index(str4,"du")
	fmt.Printf("Index:%d\n",index)

	if ret := strings.HasPrefix(str4,"http://");ret==false{
		str4 = "http://" + str4
	}
	fmt.Printf("str4 is %v\n",str4)

}

func TestForIndex()  {
	str := "hello world go"
	for index,val := range str{
		fmt.Printf("index is %d. The value is %c \n",index,val)
	}
}

func main()  {
	TestForIndex()
	StrOperation()
}