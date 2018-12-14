package main

import (
	"os"
	"log"
	"fmt"
)

func main() {
	file,err := os.Open("ifs.go")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("The file is:",file)
	file_no,errs := os.Stat("ifs.go")
	if err != nil{
		log.Fatal(errs)
	}
	fmt.Println(file_no.Size(),file_no.Name(),file_no.IsDir())
}