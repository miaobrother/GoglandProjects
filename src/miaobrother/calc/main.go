package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"strings"
)

func getInput() (string, error) { //读取用户输入的计算表达式
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}


func process(c *cli.Context) (err error) {
	for {
		express, _ := getInput()
		if len(express) == 0 {
			continue
		}
		express = strings.TrimSpace(express)
		postExpress, errRet := transPostExpress(express)
		if errRet != nil {
			err = errRet
			fmt.Println(err)
			return
		}
		result, errRet := calc(postExpress)
		if errRet != nil {
			fmt.Println("error:", errRet)
			continue
		}
		fmt.Println(result)
	}
	return
}

func calc(postExpress []string) (result int64, err error) {
	//var n1, n2 string
	var s Stack
	for i := 0; i < len(postExpress); i++ {
		var cur = postExpress[i]
		if cur[0] == '-' || cur[0] == '+' || cur[0] == '*' || cur[0] == '/' {
			n1, err := s.Pop()
			if err != nil {
				return
			}
			n2, err := s.Pop()
			if err != nil {
				return
			}
			num2, _ := strconv.Atoi(n1)
			num1, _ := strconv.Atoi(n2)
			var r1 int

			switch cur[0] {
			case '+':
				r1 = num1 + num2
			case '-':
				r1 = num1 - num2
			case '*':
				r1 = num1 * num2
			case '/':
				r1 = num1 / num2
			default:
				err = fmt.Errorf("Invalid opre")
				return
			}
			s.Push(fmt.Sprintf("%d", r1))
		} else {
			s.Push(cur)
		}
	}
	resultStr,err := s.Top()
	if err != nil{
		return
	}
	result,err = strconv.ParseInt(resultStr,10,64)
	return
}

func transPostExpress(express string) (postExpress []string, err error) {
	var opStack Stack //声明一个堆栈
	var i int
LABEL:
	for i < len(express) {
		switch {
		case express[i] >= '0' && express[i] <= '9':
			var number []byte
			for ; i < len(express); i++ {
				if express[i] < '0' || express[i] > '9' {
					break
				}
				number = append(number, express[i])
			}
			postExpress = append(postExpress, string(number))
		case express[i] == '+' || express[i] == '-' || express[i] == '*' || express[i] == '/':
			if opStack.Empty() {
				opStack.Push(fmt.Sprintf("%c", express[i]))
				i++
				continue LABEL
			}
			data, _ := opStack.Top()
			if data[0] == '(' || data[0] == ')' {
				opStack.Push(fmt.Sprintf("%c", express[i]))
				i++
				continue LABEL
			}
			if (express[i] == '+' || express[i] == '-') || (express[i] == '*' || express[i] == '/') && (data[0] == '*' || data[0] == '/') {
				postExpress = append(postExpress, data)
				opStack.Pop()
				opStack.Push(fmt.Sprintf("%c", express[i]))
				i++
				continue LABEL
			}
			opStack.Push(fmt.Sprintf("%c", express[i]))
			i++
		case express[i] == '(':
			opStack.Push(fmt.Sprintf("%c", express[i]))
			i++
		case express[i] == ')':
			for !opStack.Empty() {
				data, _ := opStack.Pop()
				if data[0] == '(' {
					break
				}
				postExpress = append(postExpress, data)
			}
			i++
		default:
			err = fmt.Errorf("invalid express:%v", express[i])
			return
		}
	}
	for !opStack.Empty() {
		data, _ := opStack.Pop()
		if data[0] == '#' {
			break
		}
		postExpress = append(postExpress, data)
	}
	return
}

func main() {
	app := cli.NewApp()
	app.Name = "calc"

	app.Usage = "calc expression"

	app.Action = func(c *cli.Context) error { // 匿名函数
		return process(c)
	}
	app.Run(os.Args)
}
