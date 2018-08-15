package main

import (
	"encoding/json"
	"fmt"
	//"github.com/gin-gonic/gin/render"
)

type UserSev struct {
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	Sex       string
}

func StructJson() (ret string, err error) {
	usersev := &UserSev{
		UserName:  "chenergou",
		FirstName: "chensanpang",
		Sex:       "女",
	}

	data, err := json.Marshal(usersev)
	if err != nil {
		err = fmt.Errorf("json marshal failed")
		return
	}
	ret = string(data)
	return
}

func test() {

	data, err := StructJson()
	if err != nil {
		fmt.Println("test struct failed,", err)
		return
	}

	var users UserSev

	err = json.Unmarshal([]byte(data), &users)
	if err != nil {
		fmt.Println("Unmarshal failed ", err)
		return
	}
	fmt.Println(users)

}

func main() {
	test()
}
