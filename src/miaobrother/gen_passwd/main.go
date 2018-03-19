package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length  int
	charset string
)

const (
	NumStr  = "0123456789"
	CharStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SpecStr = "+=@#~,.[]()!%^*$"
)

func parseArgs() {
	flag.IntVar(&length, "l", 16, "-l 是生成长度是16的密码字符集")
	flag.StringVar(&charset, "t", "num", `"-t 制定密码生成的字符串"
	"-num:只使用0-9"，"char 只使用a-z 和A-Z" ,"mix 使用数字和字母"， "advance：是使用数字，字母以及特殊字符"`)
	flag.Parse()
}

func generatePasswd() string {
	rand.Seed(time.Now().Unix())
	var passwd []byte = make([]byte, length, length)
	var sourceStr string
	if charset == "num" {
		sourceStr = NumStr
	} else if charset == "char" {
		sourceStr = CharStr
	} else if charset == "mix" {
		sourceStr = fmt.Sprintf("%s%s", NumStr, CharStr)
	} else if charset == "advance" {
		sourceStr = fmt.Sprintf("%s%s%s", NumStr, CharStr, SpecStr)
	} else {
		sourceStr = NumStr
	}
	fmt.Println("source:", sourceStr)

	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
		//fmt.Println(index)
	}

	return string(passwd)
}

func main() {
	parseArgs()
	fmt.Printf("length:%d charset:%s\n", length, charset)
	password := generatePasswd()
	fmt.Printf("The passwd is:%v\n", password)
}
