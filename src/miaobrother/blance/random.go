package main

import (
	"math/rand"
	//"fmt"
)

type Randblance struct {
}

func (r *Randblance) DoBlance(addrList []string) string {
	lenAddrList := len(addrList)
	index := rand.Intn(lenAddrList)
	return addrList[index]

}
