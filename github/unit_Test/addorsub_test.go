package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var a = 100
	var b = 200
	c := Add(a, b)
	if c != 300 {
		t.Fatal("invalid a +b")
	}
	t.Logf("testing succ")
}

func TestSub(t *testing.T) {
	var a = 100
	var b = 200
	t.Logf("a = %d b = %d\n", a, b)
	c := Sub(a, b)
	if c != -100 {
		t.Fatalf("invalid a-b ,c = %d\n", c)
	}
	t.Logf("a = %d, b = %d, c= %d", a, b, c)
}
