package main


import (
	"fmt"
	"encoding/json"
	//"unicode"
)

type Student struct {
	Id string
	Name string
	Sex string
}

type Class struct {
	Name string
	Count int
	Student []*Student
}

var rawJson = `{"Name":"111","Count":0,"Student":[{"Id":"0","Name":"stu0","Sex":"Man"},{"Id":"1","Name":"stu1","Sex":"Man"},{"Id":"2","Name":"stu2","Sex":"Man"},{"Id":"3","Name":"stu3","Sex":"Man"},{"Id":"4","Name":"stu4","Sex":"Man"},{"Id":"5","Name":"stu5","Sex":"Man"},{"Id":"6","Name":"stu6","Sex":"Man"},{"Id":"7","Name":"stu7","Sex":"Man"},{"Id":"8","Name":"stu8","Sex":"Man"},{"Id":"9","Name":"stu9","Sex":"Man"}]}`

func main() {
	c := &Class{
		Name:"111",
		Count:0,
	}
	for i := 0;i < 10; i++{
		stu := &Student{
			Name:fmt.Sprintf("stu%d",i),
			Sex : "Man",
			Id : fmt.Sprintf("%d",i),
		}
		c.Student =append(c.Student,stu)
	}
	data,err := json.Marshal(c)
	if err != nil{
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json %s\n",string(data))//Json序列化

	//Json反序列化
	fmt.Println()
	var c1 *Class = &Class{}
	json.Unmarshal([]byte(rawJson),c1)
	if err != nil{
		fmt.Println("Unmarshal failed.")
		return
	}

	//fmt.Printf("c1 %#v\n",c1)
	for _,v := range c1.Student{
		fmt.Printf("The v is %#v",v)
	}
}
