package main

import (
	"fmt"
	//"encoding/json"
	"github.com/gin-gonic/gin/json"
)

func testSlice()  {
	var M map[string]interface{}
	var s []map[string]interface{}

	M = make(map[string]interface{})

	M["username"] = "stu01"
	M["age"] = 19
	M["sex"] = "女"
	s = append(s,M)

	M["username01"] = "stu03"
	M["age01"] = 23
	M["sex01"] = "男"
	s = append(s,M)

	data,_ := json.Marshal(s)

	fmt.Printf("%s\n", string(data))
	fmt.Printf("The s is %v\n",s)
	//for i,v := range s{
	//	fmt.Printf("The i is %d v is %v\n",i,v)
	//}
}
func main()  {
	testSlice()

}