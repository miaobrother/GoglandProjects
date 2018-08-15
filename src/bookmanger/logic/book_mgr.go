package logic

import (
	"fmt"
	"sync"
)

type BookMgr struct {
	BookList       []*Book //存放
	BookStudentMap map[string][]*Student
	BookNameMap    map[string][]*Book
	BookAuthorMap  map[string][]*Book
	lock           sync.Mutex
}

func NewBookMgr() (bookMgr *BookMgr) {
	bookMgr = &BookMgr{
		BookStudentMap: make(map[string][]*Student, 10),
		BookNameMap:    make(map[string][]*Book, 10),
		BookAuthorMap:  make(map[string][]*Book, 10),
	}
	return
}

func (b *BookMgr) AddBook(book *Book) (err error) { // 增加一本书
	b.lock.Lock()
	defer b.lock.Unlock()
	b.BookList = append(b.BookList, book)
	// 更新书籍名字到同一个书籍名字对应的book列表
	bookList, ok := b.BookNameMap[book.Name]
	if !ok {
		var tmp []*Book
		tmp = append(tmp, book)
		bookList = tmp
	} else {
		bookList = append(bookList, book)
	}
	b.BookNameMap[book.Name] = bookList

	//更新书籍作者到同一个作者对应的book列表
	bookList, ok = b.BookNameMap[book.Author]
	if !ok {
		var tmp []*Book
		tmp = append(tmp, book)
		bookList = tmp
	} else {
		bookList = append(bookList, book)
	}
	b.BookNameMap[book.Author] = bookList
	return
}

func (b *BookMgr) SearchByBookName(bookName string) (bookList []*Book) { //检索
	b.lock.Lock()
	defer b.lock.Unlock()
	bookList = b.BookNameMap[bookName]
	return
}

func (b *BookMgr) SearchByBookAuthor(Author string) (bookList []*Book) { //检索
	b.lock.Lock()
	defer b.lock.Unlock()
	bookList = b.BookNameMap[Author]
	return
}

func (b *BookMgr) SearchByPublicDate(min int64, max int64) (bookList []*Book) {
	b.lock.Lock()
	defer b.lock.Unlock()
	for _, val := range bookList {
		if val.PublicDate >= min && val.PublicDate <= max {
			fmt.Printf("书籍 %v 存在与该时间段内", val)
			bookList = append(bookList, val)
			return
		}
	}
	return
}
