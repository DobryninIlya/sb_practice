package main

import "fmt"

func main() {
	var studentCount int
	var groupCount int
	var studentNum int
	fmt.Println("Студенты")
	fmt.Println("Введите число студентов:")
	fmt.Scan(&studentCount)
	fmt.Println("Введите число групп:")
	fmt.Scan(&groupCount)
	fmt.Println("Введите порядковый номер студента")
	fmt.Scan(&studentNum)
	studentInGroup := studentCount / groupCount // Число студентов в группе, с учетом того, что в каждой группе оно одинаковое
	if studentCount%groupCount != 0 {
		studentInGroup += 1 // В случае, если не делится нацело, увеличиваем число студентов группы на 1 для симметрии
	}
	numberOfGroup := studentNum / groupCount
	if studentNum%groupCount != 0 {
		numberOfGroup += 1
	}
	fmt.Println("Номер группы студента: ", numberOfGroup)
}
