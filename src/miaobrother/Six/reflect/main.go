package main

import ("reflect"
		"fmt"
)


func getTypeInfo(a interface{}){
	typeInfo := reflect.TypeOf(a)
	kind := typeInfo.Kind()
	fmt.Println("kind if a:",kind)

	num := typeInfo.NumMethod()
	fmt.Println(num)

	method,ok := typeInfo.MethodByName("SetName")
	if !ok{
		fmt.Println("It does not have method..")
	} else {
		fmt.Println(method)
	}
}

func getAllMethod(a interface{}){
	typeInfo := reflect.TypeOf(a)
	num := typeInfo.NumMethod()
	for i:= 0; i< num; i++{
		method := typeInfo.Method(i)
		fmt.Println(method)
	}
}
type Student struct {
	Name string
}

func (s *Student) SetName(name string,Age int,Sex int){
	s.Name = name
}

func (s *Student) GetName(name string)  {
	s.Name = name
}
func testGetAllMethod(){
	var stu Student
	getAllMethod(&stu)
}
func testGetTypeInfo(){
	var i int
	getTypeInfo(i)

	var stu Student
	getTypeInfo(&stu)

	var s []int
	getTypeInfo(s)

	var b [5] int
	getTypeInfo(b)
}

func testGetValueInfo()  {
	var i int
	valueinfo := reflect.ValueOf(i)
	fmt.Println(valueinfo)
}

func main()  {
	//testGetTypeInfo()
	testGetAllMethod()

}
