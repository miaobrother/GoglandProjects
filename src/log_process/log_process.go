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
	"flag"
	"github.com/influxdata/influxdb/client/v2"
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
	influxDBDsn string
}

type Message struct {
	TimeLocal                    time.Time
	ByteSent                     int
	Path, Method, Scheme, Status string
	UpstreamTime, RequestTime    float64
}

func (r *ReadFromFile) Read(rc chan []byte) {
	f, err := os.Open(r.path) // 打开文件
	if err != nil {
		panic(fmt.Sprintf("Open file error:%s\n", err.Error())) // 判断错误
	}
	defer f.Close() // 关闭文件
	f.Seek(0, 2)    // 将字符指针偏移到文件末尾
	re := bufio.NewReader(f)
	for { //循环读取文件内容
		line, err := re.ReadBytes('\n') //逐行读取内容
		if err == io.EOF {
			time.Sleep(time.Second) // 产生日志不够快 等待一秒
			continue
		} else if err != nil {
			panic(fmt.Sprintf("read from re  file error:%s\n", err.Error()))
		}
		rc <- line[:len(line)-1] //读取文件 写入到 rc channel 中 而且是循环的读

	}

}

func (w *WriteFromFile) Write(wc chan *Message) {
	infSli := strings.Split(w.influxDBDsn, "@")

	for v := range wc {
		c, err := client.NewHTTPClient(client.HTTPConfig{
			Addr:     infSli[0],
			Username: infSli[1],
			Password: infSli[2],
		})
		if err != nil {
			log.Fatal(err)
		}
		defer c.Close()

		// Create a new point batch
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  infSli[3],
			Precision: infSli[4],
		})
		if err != nil {
			log.Fatal(err)
		}

		// Create a point and add to batch
		tags := map[string]string{"Path": v.Path, "Method": v.Method, "Scheme": v.Scheme, "Status": v.Status}
		fields := map[string]interface{}{
			"UpstreamTime": v.UpstreamTime,
			"RequestTime":  v.RequestTime,
			"ByteSend":     v.ByteSent,
		}

		pt, err := client.NewPoint("access.log", tags, fields, v.TimeLocal)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)

		// Write the batch
		if err := c.Write(bp); err != nil {
			log.Fatal(err)
		}

		// Close client resources
		if err := c.Close(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("write succ..")
	}

}

//写入模块

func (l *LogProcess) ParseFromFile() { // 解析模块
	/*
		192.168.111.111 - - [14/May/2018:12:01:01 +0800] http "GET /foo?query=t HTTP/1.0" 200 2133 "-" "KeepAliveClient" "-" 1.005 1.234
	*/
	// 110.249.134.250 - - [15/May/2018:10:55:42 +0800] "GET /pcsyn/saveaisleerror?ticket=07c9fe155738429bbe606db00c3cd140&ip=192.168.111.203 HTTP/1.1" 200 36 "-" "" ""
	r := regexp.MustCompile(`([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)`)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	for v := range l.rc {
		ret := r.FindStringSubmatch(string(v))
		//fmt.Println(ret)
		//fmt.Println("The ret len is ", len(ret))
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

		byteSent, _ := strconv.Atoi(ret[7]) //发送字节
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
	var path, influxDsn string
	flag.StringVar(&path, "path", "/Users/zhangyalei/access.log", "read file path")
	flag.StringVar(&influxDsn, "influxDsn", "http://127.0.0.1:8086@root@root@db_name@s", "influx data source")
	flag.Parse()
	r := &ReadFromFile{
		path: path,
	}
	w := &WriteFromFile{
		influxDBDsn: influxDsn,
	}
	lp := &LogProcess{
		rc:    make(chan []byte), // 初始化函数
		wc:    make(chan *Message),
		read:  r,
		write: w,
	}
	go lp.read.Read(lp.rc) //读取日志
	go lp.ParseFromFile()  //解析日志
	go lp.write.Write(lp.wc)
	time.Sleep(30 * time.Second)

}
