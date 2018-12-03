// recur string ex
package main

import "fmt"

func getLenString(s string) (out string) {
	n := len(s)
	if n <= 3 {
		return s
	} else {
		return getLenString(s[:n-3]) + "," + s[n-3:]
	}
	return out

}
func main() {
	s := "wo shi ni daye "
	fmt.Println("The:", getLenString(s))

}
