package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")

	fmt.Fprintf(w, "hello...")

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	fmt.Fprintf(w, "login...")

}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/user/login", login)

	err := http.ListenAndServe(":50000", nil)

	if err != nil {
		fmt.Println("http listen failed.")
		return
	}
}
