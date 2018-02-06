package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "It is not the critic who counts;"//定义一个常量
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)//从bufio读取数据，并按照单词做分割

	for scanner.Scan(){
		fmt.Println(scanner.Text())
		fmt.Printf("%T\n",scanner.Text())//结果是字符串
	}
	if err := scanner.Err(); err != nil{
		fmt.Println(os.Stderr,"reading input:",err)
	}
}