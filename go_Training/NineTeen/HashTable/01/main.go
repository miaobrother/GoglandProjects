//package main
//
//import "fmt"
//
//func main() {
//	letter := 'A'
//
//	fmt.Println(letter)
//
//	fmt.Printf("%T \n",letter)
//}

//package main
//
//import "fmt"
//
//func main() {
//	letter := rune("A"[0])
//	fmt.Println(letter)
//}

//package main
//
//import "fmt"
//
//func main() {
//	word := "Hello"
//	letter := string((word[0]))
//	fmt.Println(letter)
//}
package main

import "fmt"

func main() {
	for i := 65;i <= 122; i++{
		fmt.Println(i,"-",string(i),"-",i%12)
	}
}
