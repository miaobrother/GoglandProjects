package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res,err := http.Get("http://www.sina.com.cn")

	if err != nil{
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	buckets := make([][]string,12)

	for scanner.Scan(){
		word := scanner.Text()
		n := hashBucket(word,12)
		buckets[n] = append(buckets[n],word)
	}

	for i :=0;i < 12; i++{
		fmt.Println(i," - ",len(buckets[i]))
	}

	fmt.Println(len(buckets))
	fmt.Println(cap(buckets))
}

func hashBucket(word string,buckets int) int{
	var sum int
	for _,v := range word{
		sum += int(v)
	}
	return sum % buckets
}

