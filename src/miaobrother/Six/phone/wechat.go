package main

import "fmt"

type WeChatPay struct {
}

func (w *WeChatPay) pay(user_id int64, money float32) error {
	fmt.Println("1.连接微信支付服务器")
	fmt.Println("2.找到对应的用户")
	fmt.Println("3.检查余额是否足够")
	fmt.Println("4.扣钱")
	fmt.Println("5.返回支付成功")
	return nil
}
