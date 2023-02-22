package storage

import "main.go/pkg/student"

type Storage interface {
	Put(student *student.Student)
	Get() map[string]*student.Student
}

type MainSotrage struct {
	M map[string]*student.Student
}

func NewStore() *MainSotrage {
	return &MainSotrage{
		M: make(map[string]*student.Student),
	}
}
func (ms *MainSotrage) Put(st *student.Student) {
	ms.M[st.Name] = st
}
func (ms *MainSotrage) Get() map[string]*student.Student {
	return ms.M
}

var Repository = make(map[string]*student.Student)
