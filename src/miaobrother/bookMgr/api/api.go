package main

import (
	"net/http"
	"miaobrother/bookMgr/logic"
	"strconv"
	"encoding/json"
)


var (
	bookMgr * logic.BookMgr
	studentMgr *logic.StudentMgr
)

func responseError(w http.ResponseWriter,code int)  {
	m := make(map[string]interface{},16)
	m["code"] = code
	m["message"] = getMessage(code)
	data,err := json.Marshal(m)
	if err != nil{
		w.Write([]byte("{\"code\":500,\"message\":\"server busy\"}"))
		return
	}
	w.Write(data)
}
func addBook(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	bookId := r.FormValue("book_id")
	name := r.FormValue("name")

	numStr := r.FormValue("num")
	author := r.FormValue("author")
	publishDateStr := r.FormValue("publish")
	num,err := strconv.Atoi(numStr)
	if err != nil{
		responseError(w,1001)
		return
	}
	book := logic.NewBook(bookId,name,num,author,publishDate)

	//queryArray := r.Form["query"]
	//data := "hello"

	w.Write([]byte(bookId))
}
func init() {
	bookMgr = logic.NewBookMgr()
	studentMgr = logic.NewStudentMgr()
}
func main() {
	http.HandleFunc("/book/add",addBook)
	http.ListenAndServe("localhost:9090",nil)
}