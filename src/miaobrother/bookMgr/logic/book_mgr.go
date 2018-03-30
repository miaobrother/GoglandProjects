package logic

import (
	//"miaobrother/Fifv/Book_Mgr/logic"
	"sync"
	//"github.com/astaxie/beego"
)

type BookMgr struct {
	BookList []*Book
	//存储bookId 到借书学生列表的信息
	BookStudentMap map[string][]*Student
	//书籍名字到书籍列表的索引
	BookNameMap map[string][]*Book
	//书籍作者到书籍列表的索引
	lock sync.Mutex
	BookAuthorMap map[string][]*Book
}

func NewBookMgr() (bookMgr*BookMgr) {
	bookMgr = &BookMgr{
		BookStudentMap:make(map[string][]*Student,16),
		BookNameMap:make(map[string][]*Book,16),
		BookAuthorMap:make(map[string][]*Book,16),
	}
	return
}

func (b *BookMgr)AddBook(book *Book)(err error)  {
	b.lock.Lock()
	defer b.lock.Unlock()

	//1.更新书籍名字到同一个书籍名字对应的book列表

	b.BookList = append(b.BookList,book)
	bookList ,ok := b.BookNameMap[book.Name]
	if !ok{
		var tmp []*Book
		tmp = append(tmp,book)
		bookList = tmp

	}else {
		bookList = append(bookList,book)
		b.BookNameMap[book.Name] = bookList
	}
	//2.更新书籍作者到同一个作者对应的book列表

	bookList ,ok := b.BookNameMap[book.Author]
	if !ok{
		var tmp []*Book
		tmp = append(tmp,book)
		bookList = tmp

	}else {
		bookList = append(bookList,book)
		b.BookNameMap[book.Author] = bookList
	}
	return
}

func (b *BookMgr) SearchByBookName(bookName string)(bookList []*Book)  {
	b.lock.Lock()
	defer b.lock.Unlock()
	bookList = b.BookNameMap[bookName]
	return
}

func (b *BookMgr) SearchByAuthor(Author string)(bookList []*Book)  {
	b.lock.Lock()
	b.lock.Unlock()
	bookList = b.BookAuthorMap[Author]
	return
}

func (b *BookMgr) SearchByPublish(min int64,max int64)(bookList []*Book)  {
	b.lock.Lock()
	b.lock.Unlock()
	for _,v := range b.BookList{
		if v.PublishDate >= min && v.PublishDate <= max{
			bookList = append(bookList,v)
		}
	}
	return
}


