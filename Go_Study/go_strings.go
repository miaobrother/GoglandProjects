package main

import (
	s "strings"
	"fmt"
)

var p = fmt.Println

func main() {
	p("Contains: ",s.Contains("test","es"))//true 因为es在test 中存在

	p("Count: ",s.Count("test","t"))// 2 字符串t在test中存在几个

	p("HasPrefix: ",s.HasPrefix("test","t")) //true 判断字符串test 是否是以t字符串开头

	p("HasSuffix: ",s.HasSuffix("test","st")) //true
	p("Index: ",s.Index("test","s"))
	p("Join: ",s.Join([]string{"a","b"}," - "))
	p("Repeat: ",s.Repeat("test",4)) //testtesttesttest
	p("Replace: ",s.Replace("foo","o","i",1))

	p("Replace: ",s.Replace("foo","o","11",1))

	p("Split: ",s.Split("a-b-c-d-","-"))
	p("Split: ",s.SplitAfter("a-b-c-d-_____","-1"))//不懂

}

