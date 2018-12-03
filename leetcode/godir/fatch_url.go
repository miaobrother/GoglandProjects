package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetUrl(urls string) {
	f, _ := os.Create("DestFile")
	resp, err := http.Get(urls)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("got a err is %v\n:", err)
		os.Exit(1)
	}
	Byte, err := io.Copy(f, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch:reading %s:%v\n", urls, err)
		os.Exit(1)
	}
	fmt.Printf("The byte is %v\n", Byte)

}
func main() {
	for _, url := range os.Args[1:] {
		if len(url) == 0 {
			log.Fatalf("The error is %v\n:", os.Stderr)
			continue
		}
		if strings.HasPrefix(url, "http://") {
			GetUrl(url)

		} else {
			url = "http://" + url
			GetUrl(url)

		}

	}

}
