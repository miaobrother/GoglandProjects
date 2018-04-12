package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"strconv"
	"regexp"
)

func sayHelloName(w http.ResponseWriter,r * http.Request)  {
	r.ParseForm()
	/*
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	*/

	for k,v := range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"hello box")
}

func login(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("method",r.Method)
	if r.Method == "GET"{
		t,_ := template.ParseFiles("login.html")
		t.Execute(w,nil)
	}else if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		phone := r.FormValue("phone")
		utype := r.FormValue("utype")
		fmt.Println(utype)

		//获取年龄转换城int
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			w.Write([]byte("数字转换异常"))
			return
		}
		if username == "" || password == "" || age == 0 {
			w.Write([]byte("username or password or age is not variable"))
			return
		}
		if age > 100 {
			w.Write([]byte("The age is too big"))
			return
		}
		if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, phone); !m {
			w.Write([]byte("phone is error"))
			return
		}
	}else {
		fmt.Println("error")
	}
}

func main() {
	//监听 走sayHomeName

	http.HandleFunc("/",sayHelloName)
	http.HandleFunc("/login",login)

	err := http.ListenAndServe(":8081",nil)
	if err != nil{
		log.Fatal("ListenAndServer:",err)
	}
}
