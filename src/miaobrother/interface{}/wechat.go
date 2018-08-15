package main

import "fmt"

type weChatPay struct {
}

func (w *weChatPay) pay(user_id int64, money float64) error {
	fmt.Println("1、连接微信支付系统中心")
	fmt.Println("2、寻找对应用户")
	fmt.Println("3、检查余额是否充足")
	fmt.Println("4、实现扣款")
	fmt.Println("5、返回支付是否成功")
	return nil

}
