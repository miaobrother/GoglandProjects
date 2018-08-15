package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func main() {
	fileName := "/Users/zhangyalei/access.tar.gz"

	insertByte := []byte("This is a test string")

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file access.log.tar.gz failed")
		return
	}
	defer file.Close()

	write := tar.NewWriter(file)

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("os stat err :%v\n", err)
	}

	hdr, err := tar.FileInfoHeader(fileInfo, "")

	if err != nil {
		fmt.Printf("tar fileInfo err %v\n", err)
	} else {
		fmt.Printf("hdr.size is %d\n", hdr.Size)
		hdr.Size = int64(len(insertByte))
	}
	err = write.WriteHeader(hdr)
	if err != nil {
		fmt.Printf("write WriteHeader err : %v\n", err)
	}

	ret, err := write.Write(insertByte)
	if err != nil {
		fmt.Printf("write ./file.tar.gz err : %v\n", err)
	} else {
		fmt.Printf("write ./file.tar.gz success . return number is %d \n", ret)
	}

	err = write.Flush()
	if err != nil {
		fmt.Printf("write flush err: %v\n", err)
	}

	err = write.Close()
	if err != nil {
		fmt.Printf("close %v failed\n", err)
	}
	writehdr, _ := os.Stat(fileName)
	fmt.Printf("the writehdr is %d\n", writehdr.Size())
}
