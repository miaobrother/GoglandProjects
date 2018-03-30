package main

import (
	"fmt"
)

func Menu()  {
	fmt.Println("1. 书籍录入:")
	fmt.Println("2. 书籍查询:")
	fmt.Println("3. 学生信息:")
	fmt.Println("4. 借书:")
	fmt.Println("5. 书籍状态:")
	fmt.Println("6. 我的:")
	fmt.Println("7. 退出:")

	var sel int
	fmt.Scanf("%d",&sel)

	switch sel {
	case 1:
		//书籍录入
	case 2:
		//数据查询
	case 3:
		//学生信息
	case 4:
		//借书
	case 5:
		//书籍状态
	case 6:
		//我的
	case 7:
		//退出


	}

}

func main() {
	Menu()
}