package main

import (
	"fmt"
	"log"
)

type Phone struct {
	PayMap map[string]Pay
}

func (p *Phone) OpenPay(name string, pay Pay) { //优化支付接口
	p.PayMap[name] = pay
}

func (p *Phone) PayMoney(name string, money float64) (err error) {
	Pay, ok := p.PayMap[name]
	if !ok {
		fmt.Println(" 支付接口不存在")
		return
	}
	err = Pay.pay(1024, 10.01)
	if err != nil {
		log.Fatalf("got error:", err)

	}
	return

}
