package main

import (
	"fmt"
	"bufio"
	"log"
	"net/http"
)

func main() {
	res,err := http.Get("https://www.baidu.com")
	if err != nil{
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)

	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)
	buckets := make([]int,12)

	for scanner.Scan(){
		n := hashBucket(scanner.Text(),12)

		buckets[n] ++
	}
	fmt.Println(buckets)
	//fmt.Println(len(buckets[65:123]))

}

func hashBucket(word string,buckets int) int  {
	var sum int
	for _,v := range word{
		sum += int(v)
	}
	return sum % buckets
}