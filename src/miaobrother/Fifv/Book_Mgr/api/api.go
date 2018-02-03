package main


import ("net/http"
	"miaobrother/Fifv/Book_Mgr/logic"
)


var (
	bookMgr *logic.BookMgr
	studentMgr *logic.StudentMgr
)

func init()  {
	bookMgr = logic.NewBookMgr()
	studentMgr = logic.NewStudentMgr()

}
func addBook(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	data := "hello"
	w.Write([]byte(data))
}


func main()  {
	http.HandleFunc("/book/add",AddBook)
	http.ListenAndServe("9090",nil)
}