package main

import "fmt"

func main() {
	var x, y int
	part := 0
	fmt.Println("Определение координатной плоскости")
	fmt.Println("Введите X и Y")
	fmt.Scan(&x)
	fmt.Scan(&y)
	if x >= 0 && y >= 0 {
		part = 1
	} else if x < 0 && y > 0 {
		part = 2
	} else if x < 0 && y < 0 {
		part = 3
	} else {
		part = 4
	}
	fmt.Println("Введенные координаты соответствуют", part, "четверти")
}
