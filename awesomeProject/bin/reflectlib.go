package bin

import (
	"fmt"
	"reflect"
)

//定义控制器函数Map类型，便于后续快捷使用
type ControllerMapsType map[string]reflect.Value

//声明控制器函数Map类型变量
var ControllerMaps ControllerMapsType

func Make() (CrMap ControllerMapsType,err error) {
	var ruTest ConnMain

	crMap := make(ControllerMapsType, 0)
	//创建反射变量，注意这里需要传入ruTest变量的地址；
	//不传入地址就只能反射Routers静态定义的方法
	vf := reflect.ValueOf(&ruTest)
	vft := vf.Type()
	//读取方法数量
	mNum := vf.NumMethod()
	//fmt.Println("NumMethod:", mNum)
	//遍历路由器的方法，并将其存入控制器映射变量中
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		//fmt.Println("index:", i, " MethodName:", mName)
		crMap[mName] = vf.Method(i) //<<<
	}
	defer func() {
		err :=recover()
		if err !=nil{
			fmt.Println(err)
		}

	}()
	return crMap,err
}