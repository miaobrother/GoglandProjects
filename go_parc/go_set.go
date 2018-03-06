package main

import (
	"fmt"
	//"strings"
	"strings"
)
//返回t在vs 中第一次出现的索引值，若没有找到t，就返回-1
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
//如果t存在于vs中，那么返回true，否则返回false
func Include(vs []string,t string) bool  {
	return Index(vs,t) >= 0
}


//如果实用vs中的任何一个字符作为函数f的参数，可以让f返回true，否则返回false
func Any(vs []string,f func(string) bool) bool  {
	for _,v := range vs{
		if f(v){
			return true
		}
	}
	return false
}


//如果分别实用vs 中所有的字符串作为f的参数都能让f返回true，那就返回true，否则false
func All(vs []string,f func(string) bool) bool  {
	for _,v := range vs{
		if !f(v){
			return false
		}
	}
	return true
}


//返回一个新的字符串切片，切片的元素为vs中所有能够让函数f返回rtue的元素
func Filter(vs []string,f func(string) bool) []string  {
	vsf := make([]string,0)
	for _,v := range vs{
		if f(v){
			vsf = append(vsf,v)
		}
	}
	return vsf
}

func Map(vs []string,f func(string) string) []string  {
	vsm := make([]string,len(vs))
	for i,v := range vs{
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	var strs = []string{"peach","apple","pear","plum"}
	fmt.Println(Index(strs,"pear"))//  res:2

	fmt.Println(Include(strs,"grape")) //  res:false

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v,"p") //  判断是否是p开头的slice

	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasSuffix(v,"m")
	}))


	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v,"e")
	}))

	fmt.Println(Map(strs,strings.ToUpper))





}
