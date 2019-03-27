package main

import (
	"fmt"
)

func main() {
	for i := 0;i < 1e5;i ++{
		httpHandler()
	}
}

func httpHandler()  {
	errch := make(chan error,1)
	resultch := make(chan int,1)
	go func() {
		defer close(errch)
		defer close(resultch)
		errch <- fmt.Errorf("shit...")
	}()
	select {
	case <- errch:
	case <- resultch:
		fmt.Println("hanpend...")


		
	}
	
}