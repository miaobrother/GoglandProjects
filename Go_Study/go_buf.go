package main

import (
	"bufio"
	"fmt"
	"os"
)

func stat(s string)  {
	var en_count int
	var sp_count int
	var int_count int
	var other_count int
	utfarr := []rune(s)
	for _,v := range utfarr{
		if v >='a' && v <= 'z'||v >='A' && v <='Z'{
			en_count++
			continue
		}
		if v == ' '{
			sp_count++
			continue
		}
		if v >= '0' && v < '9'{
			int_count++
			continue
		}
		other_count ++
	}
	fmt.Printf("en_count is %d,sp_count is %d,int_count is %d ,other is %d\n",en_count,sp_count,int_count,other_count)

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line,err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("read error ",err)
		return
	}
	fmt.Println("input:",line)

	stat(line)

}
