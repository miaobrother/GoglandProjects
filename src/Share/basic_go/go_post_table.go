package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
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
	}else {
		fmt.Println(r.PostFormValue("username"))
		fmt.Println(r.PostFormValue("password"))
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
