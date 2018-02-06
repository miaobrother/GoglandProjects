package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

)

func main() {
	res ,err := http.Get("http://www.gutenberg.org")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("以下是...")
	bs ,err := ioutil.ReadAll(res.Body)

	res.Body.Close()

	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%s",bs)
}