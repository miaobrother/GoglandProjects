package main

import "fmt"

func main() {
	runefanzhuan()
}

func runefanzhuan() {
	var str = "hello中文"
	var res string
	var r = []rune(str)
	for i := 0; i < len(r); i++ {
		res = res + fmt.Sprintf("%c", r[len(r)-1-i])

	}
	fmt.Println(res)

}
