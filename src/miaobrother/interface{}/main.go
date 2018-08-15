package main

import (
	"fmt"
	"log"
)

func main() {
	phone := &Phone{
		PayMap: make(map[string]Pay, 10),
	}
	//phone.OpenPayAliPay()//注册阿里支付功能
	//phone.OpenPay("ali_pay",&AliPay{})
	phone.OpenPay("wechat_pay", &weChatPay{})
	err := phone.PayMoney("wechat_pay", 1.01)
	if err != nil {
		log.Fatal("wechat pay failed.error is:", err)
		return
	}
	fmt.Println("pay success")

}
