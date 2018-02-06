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
	buckets := make([]int,200)

	for scanner.Scan(){
		n := hashBucket(scanner.Text())

		buckets[n] ++
	}
	fmt.Println(buckets[65:123])
	//fmt.Println(len(buckets[65:123]))

}

func hashBucket(word string) int  {
	return int(word[0])
}