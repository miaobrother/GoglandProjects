package main

import (
	"fmt"
	"reflect"
)

type Teac struct {
	Name string
}

func main() {
	testStruct()

}
func (t *Teac) Set(name string) string {
	nameNew := t.Set("nimei")
	return nameNew

}
func (t *Teac) GetName(name string) {
	t.Name = name
}

func testStruct() {
	var t Teac
	valueInfo := reflect.ValueOf(&t)
	filedNum := valueInfo.Elem().NumField()
	fmt.Println(filedNum)

}
