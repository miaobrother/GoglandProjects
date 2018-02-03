package main

import ("fmt"
		"math/rand"
		//"strings"
)

func test()  {
	var a [10] int
	for i := 0; i < len(a); i++{
		fmt.Printf("%p\n",&a[i])
	}

	for index,_ := range a {
		fmt.Printf("a[%d] = %d\n",index,a[index])
	}
}

func test1()  {
	var a [5] int = [5] int {1,2,3,4,5} //this is a array
	var b [5] int // this is a list
	b = a
	fmt.Printf("%v\n",b)
	b[0] = 200
	fmt.Printf("%v\n %v\n",b,a)
}

func test2()  {
	var a[5] int = [5] int {1,1,1,1}
	fmt.Printf("%v\n",a)
	var b = [...] int {1,2,3,4,5}
	fmt.Printf("%v\n",b)

	var d [5]string = [5]string{1:"zhang",4:"lisi"}
	fmt.Printf("%#v\n",d)
}

func test4()  {

	var a [4][2]int
	for i := 0;i<4; i++{
		for j :=0;j<2;j++{
			a[i][j] = i*j
		}
	}
	for i := 0;i<4;i++{
		for j := 0;j<2;j++{
			fmt.Printf("%d",a[i][j])
		}
		fmt.Println()
	}

}

func genRand()  {
	var a  [100]int
	for i := 0; i< len(a); i++{
		//赋值
		a[i] = rand.Int()
		//fmt.Printf("%d\n",a[i])
	}
	for i := 0; i<len(a);i++{
		//获取下表=i的元素的值
		fmt.Printf("%d\n",a[i])
	}
}

func genRandStr()  {
	var a [100]string
	var b string = "abcdefghijklmnopqrstuvwxyz"
	//var runeArr = []rune(b)
	for i := 0; i<len(a);i++{
		var str string
		for j :=0;j <32;j ++{
			index := rand.Intn(len(b))
			//c:= b[index]
			str = str + fmt.Sprintf("%c",b[index])
		}
		a[i] = str
		fmt.Printf("a[%d] = %s\n",i,a[i])
	}
}
func main()  {
	//test()
	//test1()
	//test2()
	//test4()
	//genRand()
	genRandStr()
}
