package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	//"os"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	ch := make(chan string)

	for _, url2 := range os.Args[1:] {
		go fetch(url2, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url2 string, ch chan<- string) {
	start1 := time.Now()
	resp, err := http.Get(url2)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url2, err)
		return
	}

	secs := time.Since(start1).Seconds()
	ch <- fmt.Sprintf("%.2fs %d %s", secs, nbytes, url2)
}
