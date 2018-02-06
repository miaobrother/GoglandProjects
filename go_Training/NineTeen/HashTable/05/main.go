package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res,err := http.Get("http://www.douban.com")
	if err != nil{
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)

	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}