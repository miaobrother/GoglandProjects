package logic

import (
	"sync"
	"fmt"
)

type StudentMgr struct {
	Student map[int]*Student
	lock sync.Mutex
}

func NewStudentMgr()(*StudentMgr)  {
	return &StudentMgr{}
}

func (s *StudentMgr)AddStudent(stu *Student)  {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.StudentMap[stu.Id] = stu
}

func (s * StudentMgr)GetStudentById(id int)(stu *Student,err error)  {
	s.lock.Lock()
	defer s.lock.Unlock()

	stu,ok := s.StudentMap[id]
	if !ok{
		err = fmt.Errorf("student id %d is not exists!",id)
		return
	}
	return
}

func (s *StudentMgr)GetStudentBorrowsBook (id int) (bookList []*Book,err error)  {
	stu ,err := s.GetStudentById(id)
	if err != nil{
		return
	}
	bookList = stu.GetBookList()
	return
}
