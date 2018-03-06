package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	dateCmd := exec.Command("date")
	dateOut,err := dateCmd.Output()
	if err != nil{
		panic(err)
	}
	fmt.Println(">date:")
	fmt.Println(string(dateOut))//Tue Mar  6 15:25:16 CST 2018

	grepCmd := exec.Command("grep ","Hello")

	grepIn,_ := grepCmd.StdinPipe()

	grepOut,_ := grepCmd.StdoutPipe()

	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes,_:= ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println(">grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash","-c","ls -a -l -h")
	lsOut,err := lsCmd.Output()
	if err != nil{
		panic(err)
	}
	fmt.Println(">ls -a -l -h")
	fmt.Println(string(lsOut))
}