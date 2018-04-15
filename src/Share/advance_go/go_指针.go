package main

import . "fmt"

func main() {
	var in int = 100
	Printf("in = %d\n",in)
	Printf("&in = %p\n",&in) // 变量的内存

	//保存某个变量的地址，需要指针类型   *int 保存的int的地址   **int 保存的是 *int地址

	var pp *int//定义一个变量pp  类型为*int
	pp = &in
	Printf("The pp is %v\n",*pp)

	*pp = 777 //*pp操作的不是pp的 内存， 而是操作的pp所指向的内存 就是in
	Printf("The in is %v\n",in)
	Printf("The pp is %v\n",*pp)
}