package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

func NewStudent() *Student {
	return &Student{}
}

type Storage interface {
	Put(*Student)
	Get() map[string]*Student
}

type MainSotrage struct {
	M map[string]*Student
}

func NewStore() *MainSotrage {
	return &MainSotrage{
		M: make(map[string]*Student),
	}
}
func (ms *MainSotrage) Put(st *Student) {
	ms.M[st.Name] = st
}
func (ms *MainSotrage) Get() map[string]*Student {
	return ms.M
}

type App struct {
	repository Storage
}

func NewApp(repository Storage) *App {
	return &App{
		repository: repository,
	}
}

func (a *App) run() {
	for {
		fmt.Println("Введите Имя, возраст и оценку. Для выхода: exit")
		var input string
		fmt.Scan(&input)
		if input != "exit" {
			student := NewStudent()
			student.Name = input
			fmt.Scan(&input)
			student.Age, _ = strconv.Atoi(input)
			fmt.Scan(&input)
			student.Grade, _ = strconv.Atoi(input)
			a.repository.Put(student)
		} else {
			fmt.Println(a.studentNameList())
			fmt.Println("Done")
			return
		}
	}
}

func (a *App) studentNameList() string {
	var result string
	fmt.Println(a.repository.Get())
	for k, v := range a.repository.Get() {
		result += string(k) + " " + strconv.Itoa(v.Age) + " " + strconv.Itoa(v.Grade)
	}
	return result
}

func main() {
	repository := NewStore()

	var app = NewApp(repository)
	app.run()

}
