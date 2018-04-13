package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64 = 0

	for i :=0;i < 50 ;i ++{
		go func() {
			for{
				atomic.AddUint64(&ops,1)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)

	//获取结果
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:",opsFinal)

	//
	var tmpOps uint64 = 0

	t1 := time.Now().UnixNano()
	for tmpOps < 1e7{
		go func() {
			atomic.AddUint64(&tmpOps,1)
		}()
	}
	t2 := time.Now().UnixNano()
	t := t2-t1
	fmt.Printf("The pro cost is :%v ms\n",t/1e6)
	fmt.Println(tmpOps)
}
