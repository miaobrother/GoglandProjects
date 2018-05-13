package main

import (
	"testing"
)

func BenchmarkAdd(b *testing.B) {//基准测试  或叫 压力测试
	for i := 0;i < b.N;i ++{
		a := 10
		b := 20
		Add(a,b)

	}
}

func BenchmarkSub(b *testing.B) {
	for i := 0;i <b.N;i++{
		a := 10
		b := 20
		Sub(a,b)
	}
}