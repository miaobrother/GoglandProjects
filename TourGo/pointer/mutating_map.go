package main

import (
	"fmt"
	//"os/exec"
)

func main() {
	Ma := make(map[string]string)
	fmt.Println(Ma)
	Ma["name"] = "zhangSan"
	fmt.Println("The value:",Ma["name"])

	Ma["Sex"] = "Man"
	fmt.Println("The value:",Ma["Sex"])

	delete(Ma,"Sex")//删除Map 中的key

	fmt.Println(Ma)
	_,ok := Ma["namae"]
	if ok{
		fmt.Println("  存在。")
	}else {
		fmt.Println("不存在。")
	}

}
