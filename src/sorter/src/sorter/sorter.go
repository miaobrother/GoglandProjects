package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)
	for {
		line, isPrefix, errOne := br.ReadLine()
		if errOne != nil {
			if errOne != io.EOF {
				err = errOne
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected")
			return
		}
		str := string(line)

		strVal, errOne := strconv.Atoi(str)
		if errOne != nil {
			err = errOne
			return
		}
		values = append(values, strVal)
	}
	return
}
func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("Read values:", values)
	} else {
		fmt.Println(err)
	}
}
