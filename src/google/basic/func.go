package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("invailed string")
	}
}

func div(a, b int) (q, r int) {

	//fmt.Printf("The a is %d b is %d\n", a, b)
	return a / b, a % b
}

func apply(oc func(int, int) int, a, b int) int {
	p := reflect.ValueOf(oc).Pointer()
	ocName := runtime.FuncForPC(p).Name()
	fmt.Printf("The oc type is %v\n", p)
	fmt.Printf("Calling function %s with args "+"(%d,%d)", ocName, a, b)
	return oc(a, b)

}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int { //可变参数
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b *int) { //指针传递
	*a, *b = *b, *a
}
func main() {
	fmt.Println(div(13, 3))
	fmt.Println("operation is:", eval(1, 5, "*"))
	fmt.Println(apply(pow, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	var a, b = 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
