package main

import (
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	//"golang.org/x/text/encoding/simplifiedchinese"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
)

func main() {

	resp, err := http.Get("http://www.fang.com/SoufunFamily.htm")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
	}
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder()) //GbK 转化至 utf8
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(all))
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
