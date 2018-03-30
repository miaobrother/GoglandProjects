package main

import (
	"fmt"
	"math/rand"
	"time"
	//"expvar"
	"sort"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //生成随机种子

	var m map[string]int = make(map[string]int, 1024) //初始化一个Map

	for i := 0; i < 256; i++ {
		key := fmt.Sprintf("stud%d", i)
		value := rand.Intn(256)
		m[key] = value
	}
	var keys []string = make([]string, 0, 257)
	for k, v := range m {
		fmt.Printf("The key is:%v,The value is:%v\n", k, v)
		keys = append(keys, k) //输出是非排序好的，
	}
	//fmt.Println(len(m))
	sort.Strings(keys) //使用sort函数进行排序
	for _, val := range keys {
		fmt.Printf("key :%s val is %d\n", val, m[val])
	}
}
