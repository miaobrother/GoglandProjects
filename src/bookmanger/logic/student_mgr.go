package logic

import (
	"fmt"
	"sync"
)

type StudentMgr struct {
	StudentMap map[int]*Student
	lock       sync.Mutex
}

func NewStudentMgr() *StudentMgr {
	return &StudentMgr{}
}

func (s *StudentMgr) AddStudent(student *Student) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.StudentMap[student.Id] = student
}

func (s *StudentMgr) SearchStudentById(id int) (stu *Student, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	stu, ok := s.StudentMap[stu.Id]
	if !ok {
		err = fmt.Errorf("student id %d is not exist.", id)
		return
	}
	return
}

func (s StudentMgr) GetStudentBorrowsBook(id int) (bookList []*Book, err error) {
	stu, err := s.SearchStudentById(id)
	if err != nil {
		return
	}
	bookList = stu.GetBookList()
	return
}
