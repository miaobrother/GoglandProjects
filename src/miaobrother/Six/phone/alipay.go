package main

import "fmt"

type AliPay struct {
}

func (w *AliPay) pay(user_id int64, money float32) error {
	fmt.Println("1.连接支付宝服务器")
	fmt.Println("2.找到对应的用户")
	fmt.Println("4.扣钱")
	fmt.Println("5.返回支付成功")
	return nil
}
