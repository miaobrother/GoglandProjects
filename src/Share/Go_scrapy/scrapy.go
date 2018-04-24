package main

import (
	"fmt"
	//"net/http"
	"log"
	"net/http"
	"os"
	"strconv"
	"io"
)

func SpiderPage(i int, page chan int) {
	url := "https://www.v2ex.com/go/python?p=" + strconv.Itoa((i)*1)
	//fmt.Printf("The scrapy url is %s\n",url)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err", err)
		return
	}
	//把内容写入文件
	fileName := strconv.Itoa(i) + ".html"
	f, err1 := os.Create(fileName)
	if err1 != nil {
		log.Fatal(err1)
		return
	}
	f.WriteString(result)
	f.Close()
	page <- i
}
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		fmt.Println("The http.Get err is", err1)
		return
	}

	defer resp.Body.Close()
	//读取数据

	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err = ", err)
			break
		}
		result += string(buf[:n])
	}
	return

}
func DoScrapy(start, end int) {
	page := make(chan int)
	fmt.Printf("爬虫工作开始 从%d 到%d\n", start, end)
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)

	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面抓去完成\n", <-page)
	}

}
func main() {
	var start, end int
	fmt.Println("please input start page:")
	fmt.Scan(&start)
	fmt.Println("please input end page:")
	fmt.Scan(&end)
	DoScrapy(start, end)
}
