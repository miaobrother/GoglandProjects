//echo return command-line arg..
package main

import (
	"fmt"
	"os"
	//"strings"
	"strings"
	"time"
)

func main() {
	/*
		start_t := time.Now()
		var s,step string
		//var tmp []string
		step = " "
		for _,val := range os.Args[1:]{
			s += step + val

		}
		//tmp = append(tmp,s)

		stop_t := time.Now()
		fmt.Println(stop_t.Sub(start_t))
	*/

	start_t := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	stop_t := time.Now()
	fmt.Println(stop_t.Sub(start_t))
}
