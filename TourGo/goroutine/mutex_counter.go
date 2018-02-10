package main

import (
	"fmt"
	"time"
	"sync"
	//"crypto/aes"
)

type safeCounter struct {//定义一个safeCounter 类，并且 传入v 是map类型，key 是字符串，value是int
	v map[string] int
	mux sync.Mutex//定义一把锁  别名是mux
}

func (s *safeCounter) Inc(key string)  {//增加给定key的计数器的值
	s.mux.Lock()
	s.v[key]++ //lock之后智能有一个goroutinue访问s.v
	s.mux.Unlock()

}

func (s *safeCounter) Value(key string)int  {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.v[key]
}

func main() {
	s := safeCounter{v:make(map[string]int)}
	 for i:=0;i < 100; i++{
	 	go s.Inc("somekey")
	 }

	 time.Sleep(time.Second)

	 fmt.Println(s.Value("somekey"))
}