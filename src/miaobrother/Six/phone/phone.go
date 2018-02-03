package main

import "fmt"

type Phone struct {
	PayMap map[string]Pay
}

func (p *Phone) OpenWeChatPay() {
	weChatPay := &WeChatPay{}

	p.PayMap["wechat_pay"] = weChatPay
}

func (p *Phone) OpenAliPay() {
	aliPay := &AliPay{}
	p.PayMap["ali_pay"] = aliPay
}

func (p *Phone) PayMoney(name string, money float32) (err error) {
	pay, ok := p.PayMap[name]
	if !ok {
		err = fmt.Errorf("It does not support %s", name)
		return
	}
	err = pay.pay(1024, 8888)
	return
}
