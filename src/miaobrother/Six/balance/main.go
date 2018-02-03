package main

import (
	"fmt"
	"math/rand"
)


func doBalance(balance Balance,addrList []string )(addr string) {
	return balance.DoBalance(addrList)
}

func main()  {
	var addrList []string
	for i:= 0 ;i <10; i++ {
		addr := fmt.Sprintf("%d.%d.%d.%d",rand.Intn(255),rand.Intn(255),rand.Intn(255),rand.Intn(255))
	addrList = append(addrList,addr)

	}
	RandBalance := &RandBalance{}
	for i:= 0; i<10;i++{
		addr := doBalance(RandBalance,addrList)
		fmt.Println(addr)

	}

}