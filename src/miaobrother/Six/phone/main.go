package main

import "fmt"

func main() {
	phone := &Phone{
		PayMap: make(map[string]Pay, 19),
	}
	phone.OpenAliPay()
	err := phone.PayMoney("wechat_Pay", 20.20)
	if err != nil {
		fmt.Printf("Pay failed..The reson is %v\n", err)
		fmt.Println("Please Use Alipay")
		err = phone.PayMoney("ali_pay", 20.22)
		if err != nil {
			fmt.Printf("The Pay failed.%v", err)
		}
		return
	}
	fmt.Println("The Pay success Welcome to you..")
}
