package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
			b, err := ioutil.ReadAll(resp.Body)
			fmt.Println(resp.Status)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			}
			fmt.Printf("%s\n", b)

		} else {
			fmt.Println("是否需要添加http://")
			url = "http://" + url
			fmt.Println(url)
		}

	}
}
