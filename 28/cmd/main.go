package main

import (
	"fmt"
	"main.go/pkg/storage"
	"main.go/pkg/student"
	"strconv"
)

type App struct {
	repository storage.Storage
}

func NewApp(repository storage.Storage) *App {
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
			st := student.NewStudent()
			st.Name = input
			fmt.Scan(&input)
			st.Age, _ = strconv.Atoi(input)
			fmt.Scan(&input)
			st.Grade, _ = strconv.Atoi(input)
			a.repository.Put(st)
		} else {
			fmt.Println(a.studentNameList())
			fmt.Println("Done")
			return
		}
	}
}

func (a *App) studentNameList() string {
	var result string
	for k, v := range a.repository.Get() {
		result += string(k) + " " + strconv.Itoa(v.Age) + " " + strconv.Itoa(v.Grade) + "\n"
	}
	return result
}

func main() {
	repository := storage.NewStore()
	var app = NewApp(repository)
	app.run()

}
