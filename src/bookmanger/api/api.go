package main

import (
	"bookmanger/logic"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	bookMgr    *logic.BookMgr
	studentMgr *logic.StudentMgr
)

func ResponseError(w http.ResponseWriter, code int) {
	m := make(map[string]interface{}, 10)
	m["code"] = code
	m["message"] = GetMessage(code)

	data, err := json.Marshal(m)
	if err != nil {
		w.Write([]byte("{\"code\":500,}"))
		log.Fatal("get json failed.", err)
		return
	}
	w.Write(data)
	fmt.Println(string(data))
}
func ResponseSuccess(w http.ResponseWriter, code int, data interface{}) {
	m := make(map[string]interface{}, 10)
	m["code"] = code
	m["message"] = GetMessage(code)
	m["data"] = data

	dataByte, err := json.Marshal(m)
	if err != nil {
		w.Write([]byte("{\"code\":500,}"))
		log.Fatal("get json failed.", err)
		return
	}
	fmt.Println(string(dataByte))
}

func init() {
	bookMgr = logic.NewBookMgr()
	studentMgr = logic.NewStudentMgr()
}
func AddBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bookId := r.FormValue("bookId")
	name := r.FormValue("name")
	numS := r.FormValue("num")
	author := r.FormValue("author")
	publicDate := r.FormValue("public")
	num, err := strconv.Atoi(numS)
	if err != nil {
		ResponseError(w, ErrInvalidNum)
		return

	}
	intI, err := strconv.Atoi(publicDate)
	if err != nil {
		log.Fatal("get error:", err)
		return
	}

	book := logic.NewBook(bookId, name, num, author, int64(intI))
	fmt.Println("The book is:", book)
	err = bookMgr.AddBook(book)
	if err != nil {
		ResponseError(w, ErrInvalidNum)
		return
	}
	fmt.Println("add book success")
	ResponseSuccess(w, ErrSuccess, book)

}

func SearchByBookName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	if len(name) == 0 {
		fmt.Println("got a err:")
	}
	err := bookMgr.SearchByBookName(name)
	for _, val := range err {
		fmt.Println("Get book is :", val)
	}
	fmt.Println("search book by name success")
	ResponseSuccess(w, ErrSuccess, err)

}

func SearchByBookAuthor(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	author := r.FormValue("author")
	if len(author) == 0 {
		fmt.Println("got a err:")
	}
	err := bookMgr.SearchByBookAuthor(author)
	for _, val := range err {
		fmt.Println("Get book is :", val)
	}
	fmt.Println("search book by author success")
	ResponseSuccess(w, ErrSuccess, err)

}

func main() {
	http.HandleFunc("/book/add", AddBook)
	http.HandleFunc("/book/search", SearchByBookName)
	http.HandleFunc("/book/searchByauthor", SearchByBookAuthor)
	http.ListenAndServe(":8080", nil)
}
