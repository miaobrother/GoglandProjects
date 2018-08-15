package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var urls = []string{
	"http://www.pasq.com",
	"http://www.google.com",
	"http://www.taobao.com",
}

func main() {
	for _, v := range urls {
		c := http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					timeout := time.Second * 2
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}
		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head  %s failed,err :%v\n", v, err)
			continue
		}
		fmt.Printf("head %s successful,status:%v\n", v, resp.Status)
	}

}
