package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//"net/url"
)

func main() {
	for _, url1 := range os.Args[1:] {
		resp, err := http.Get(url1)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fatch:%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url1, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
