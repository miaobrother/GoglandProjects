package logic

import (
	"errors"
	"sync"
	//"time"
)

type Book struct {
	BookId     string
	Name       string
	Num        int
	Author     string
	PublicDate int64
	lock       sync.Mutex //互斥锁
}

func NewBook(bookId, name string, num int, author string, publicDate int64) (book *Book) {
	book = &Book{
		BookId:     bookId,
		Name:       name,
		Num:        num,
		Author:     author,
		PublicDate: publicDate,
	}
	return

}

func (b *Book) Borrow() (err error) { //借书方法
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.Num <= 0 { //判断书籍
		err = errors.New("book not enghout..")
		return
	}
	b.Num = b.Num - 1
	return
}

func (b *Book) PayBack() (err error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.Num = b.Num + 1
	return
}
