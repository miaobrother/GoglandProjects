package logic

import (
	"sync"

)

type BookMgr struct {
	BookList []*int
	//存储bookId到借书学生列表
	BookStudentMap map[string][]*Student
	//书籍名字到书籍列表的索引
	BookNameMap map[string][]*Book
	//书籍作者到书籍列表的索引
	BookAuthorMap map[string][]*Book
	lock          sync.Mutex
}

func NewBookMgr() (BookMgr *BookMgr) {
	BookMgr = &BookMgr{
		BookStudentMap: make(map[string][]*Student, 16),
		BookNameMap:    make(map[string][]*Book, 16),
		BookAuthorMap:  make(map[string][]*Book, 16),
	}
	return
}

func (b *BookMgr) AddBook(book *Book) (err error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.BookList = append(b.BookList, book)
	booklist, ok := b.BookNameMap[book.Name]

	if !ok {
		var tmp []*Book
		tmp = append(tmp, book)
		booklist = tmp

	} else {
		booklist = append(booklist, book)
		b.BookNameMap[book.Name] = booklist
	}
	b.BookNameMap[book.Name] = booklist
	return
}

func (b *BookMgr) SearchByBookName(bookName string) (booklist []*Book) {
	b.lock.Lock()
	defer b.lock.Unlock()
	booklist = b.BookNameMap[bookName]
	return
}

func (b *BookMgr) SearchByPublish(min int64, max int64) (booklist []*Book) {
	b.lock.Lock()
	defer b.lock.Unlock()
	for _, v := range b.BookList {
		if v.PublishDate > min && v.PublishDate <= max {
			booklist = append(booklist, v)
		}
	}
	return
}
