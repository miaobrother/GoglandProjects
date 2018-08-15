package main

import "fmt"

type AliPay struct {
}

func (ali *AliPay) pay(user_id int64, money float64) error {
	fmt.Println("1、连接阿里支付系统中心")
	fmt.Println("2、寻找对应用户")
	fmt.Println("3、实现扣款")
	fmt.Println("4、返回支付是否成功")
	return nil

}
