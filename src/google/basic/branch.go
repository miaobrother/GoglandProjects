package basic

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "/Users/zhangyalei/access.log"

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
