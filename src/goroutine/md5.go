package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	str := "abc123"

	data := []byte(str)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)
	fmt.Println(md5Str)

	w := md5.New()
	io.WriteString(w, str)
	bw := w.Sum(nil)
	md5StrTwo := hex.EncodeToString(bw)
	fmt.Println(md5StrTwo)
}
