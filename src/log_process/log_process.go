package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	//"log"
	"bufio"
	"io"
	//"regexp/syntax"
	"log"
	"net/url"
	"regexp"
	"strconv"
)

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan *Message)
}

type LogProcess struct {
	rc    chan []byte   // 读取chan
	wc    chan *Message // 写入chan
	read  Reader
	write Writer
}

type ReadFromFile struct {
	path string
}
type WriteFromFile struct {
	influxbSource string
}

type Message struct {
	TimeLocal                    time.Time
	ByteSent                     int
	Path, Method, Scheme, Status string
	UpstreamTime, RequestTime    float64
}

func (r *ReadFromFile) Read(rc chan []byte) {
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("Open file error:%s\n", err.Error()))
	}
	defer f.Close()
	f.Seek(0, 2)
	re := bufio.NewReader(f)
	for {
		line, err := re.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("read from re  file error:%s\n", err.Error()))
		}
		rc <- line //读取文件 写入到 rc channel 中 而且是循环的读

	}

}

func (w *WriteFromFile) Write(wc chan *Message) {
	for {
		fmt.Println(<-wc)

	}
	//写入模块

}

//110.228.7.45 - - [14/May/2018:12:01:01 +0800] "GET /pcsyn/saveaisleerror?ticket=e1e5a602061111e683d6fcaa14f0269f&ip=192.168.111.202 HTTP/1.1" 502 166 "-" "Java/1.8.0_65" "-"
func (l *LogProcess) ParseFromFile() { // 解析模块
	/*
		192.168.111.111 - - [14/May/2018:12:01:01 +0800] http "GET /foo?query=t HTTP/1.0" 200 2133 "-" "KeepAliveClient" "-" 1.005 1.234
		([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"
		(.*?)\"\s+\"([\d\.-]+)\"s+([\d\.-]+)\s+([\d\.-]+)
	*/

	r := regexp.MustCompile(`([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)`)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	for v := range l.rc {
		ret := r.FindStringSubmatch(string(v))
		//fmt.Println("The ret len is ",len(ret))
		if len(ret) != 14 {
			log.Println("FindStringSubmatch fail", string(v))
			continue
		}

		t, err := time.ParseInLocation("02/Jan/2006:15:04:05 +0800", ret[4], loc)
		if err != nil {
			log.Println("parse time failed")

		}
		message := &Message{}
		message.TimeLocal = t

		byteSent, _ := strconv.Atoi(ret[8])
		message.ByteSent = byteSent

		reqSli := strings.Split(ret[6], " ")
		if len(reqSli) != 3 {
			log.Println("strings split failed", ret[6])
		}
		message.Method = reqSli[0]

		u, err := url.Parse(reqSli[1])
		if err != nil {
			log.Println("url failed", err)
			continue
		}
		message.Path = u.Path
		message.Scheme = ret[5]
		message.Status = ret[7]
		upstreamTime, _ := strconv.ParseFloat(ret[12], 64)
		costTime, _ := strconv.ParseFloat(ret[13], 64)
		message.UpstreamTime = upstreamTime
		message.RequestTime = costTime
		l.wc <- message
	}

}

func main() {
	r := &ReadFromFile{
		path: "/Users/zhangyalei/access.log",
	}
	w := &WriteFromFile{
		influxbSource: "/Users/zhangyalei/123.log",
	}
	lp := &LogProcess{
		rc:    make(chan []byte),
		wc:    make(chan *Message),
		read:  r,
		write: w,
	}
	go lp.read.Read(lp.rc)
	go lp.ParseFromFile()
	go lp.write.Write(lp.wc)
	time.Sleep(30 * time.Second)

}
