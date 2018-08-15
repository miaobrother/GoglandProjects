/*
|---miaobrother
|   |----arithmetic
|   |   |----choice.go
|   |   |----insert_Ari.go
|   |   |----leetcode_Sum.go
|   |   |----sum_array.go
|   |----bookMgr
|   |   |----api
|   |   |   |----api.go
|   |   |   |----error_code.go
|   |   |----logic
|   |   |   |----book.go
|   |   |   |----book_mgr.go
|   |   |   |----student.go
|   |   |   |----student_mgr.go
|   |----calc
|   |   |----main.go
|   |   |----stack.go
|   |----coins
|   |   |----main.go
|   |----datetime
|   |   |----main.go
|   |----file_read
|   |   |----change_String.go
|   |   |----main.go
|   |   |----reverse_str.go
|   |----for_para
|   |   |----main.go
|   |----gen_passwd
|   |   |----gen_passwd
|   |   |----main.go
|   |----go_io
|   |   |----Args.go
|   |   |----bufio.go
|   |   |----cat.go
|   |   |----copyFile.go
|   |   |----defer.go
|   |   |----flag.go
|   |   |----input.go
|   |   |----ioutil.go
|   |   |----main.go
|   |   |----read_gzip.go
|   |   |----source.txt
|   |   |----target.txt
|   |   |----test.go
|   |   |----urfave.go
|   |   |----write.go
|   |   |----writeFile.go
|   |----go_os
|   |   |----main.go
|   |----goroutine
|   |   |----main.go
|   |----json
|   |   |----main.go
|   |----main.go
|   |----map
|   |   |----change_map.go
|   |   |----count_string.go
|   |   |----nil_map.go
|   |   |----rand_map.go
|   |   |----sliceMap.go
|   |   |----student_info.go
|   |----multiPlication
|   |   |----main.go
|   |----pointer
|   |   |----main.go
|   |----prime
|   |   |----goNarcissistic.go
|   |   |----go_count_string.go
|   |   |----main.go
|   |----string
|   |   |----main.go
|   |----struct_para
|   |   |----go_inherit.go
|   |   |----go_method.go
|   |   |----go_struct.go
|   |   |----go_struct_mem.go
|   |   |----go_struct_point.go
|   |   |----go_struct_tag.go
|   |----tree
|   |   |----tree
|   |   |----tree.go
*/
package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ListDir(dirPath string, deep int) (err error) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	if deep == 1 {
		fmt.Printf("|---%s\n", filepath.Base(dirPath))
	}

	sep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() {
			fmt.Printf("|")
			for i := 0; i < deep; i++ {
				fmt.Printf("   |")
			}
			fmt.Printf("----%s\n", fi.Name())
			ListDir(dirPath+sep+fi.Name(), deep+1)
			continue
		}
		fmt.Printf("|")
		for i := 0; i < deep; i++ {
			fmt.Printf("   |")
		}
		fmt.Printf("----%s\n", fi.Name())
	}
	return
}

func main() {
	app := cli.NewApp()
	app.Name = "tree"
	app.Usage = "list all file"

	app.Action = func(c *cli.Context) error {
		var dir string = "."
		if c.NArg() > 0 {
			dir = c.Args()[0]
		}
		ListDir(dir, 1)
		return nil
	}
	app.Run(os.Args)
}
