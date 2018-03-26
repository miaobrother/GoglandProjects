package main

import (
	"bufio"
	"io"
	"fmt"
	"os"
	"flag"
)

func cat(r *bufio.Reader)  {
	for{
		buf,err := r.ReadBytes('\n')
		if err == io.EOF{
			break
		}
		fmt.Fprintf(os.Stdout,"%s",buf)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0{
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0;i < flag.NArg();i++{
		f,err := os.Open(flag.Arg(i))
		if err != nil{
			fmt.Fprintf(os.Stderr,"%s :error reading from %s\n",os.Args[0],flag.Args(),err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}