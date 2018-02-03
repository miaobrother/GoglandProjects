package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var sum int
	sum = Add(5, 6)
	if sum != 11 {
		t.Fatal("add is no right, sum%v expcted:11", sum)
	}

	t.Log("add is ok")
}
