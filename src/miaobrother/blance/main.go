package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//import "fmt"

func doBlance(blance Blance, addrList []string) (addr string) {
	return blance.DoBlance(addrList)

}

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	var addrList []string
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter num:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatal("got a error:", err)
		return
	}
	inputSplit := strings.Replace(input, "\n", "", -1)
	inputInt, err := strconv.Atoi(inputSplit)
	if err != nil {
		log.Fatal("strconv failed:", err)
		return
	}
	if inputInt == 1 {
		randBlance := &Randblance{}
		for i := 0; i < 10; i++ {
			addr := fmt.Sprintf("%d.%d.%d.%d:8080", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
			addrList = append(addrList, addr)
			addrL := doBlance(randBlance, addrList)
			fmt.Println(addrL)
		}
	} else {
		roundblance := &Roundblance{}
		for i := 0; i < 5; i++ {
			addr := fmt.Sprintf("%d.%d.%d.%d:8080", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
			addrList = append(addrList, addr)
			addrL := doBlance(roundblance, addrList)
			fmt.Println(addrL)
		}

	}

}
