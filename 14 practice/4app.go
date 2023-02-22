package main

import "fmt"

var variable1 = 10
var variable2 = 200
var variable3 = 50

func someFunc1(n int) int {
	return n + variable1
}
func someFunc2(n int) int {
	return n + variable2
}
func someFunc3(n int) int {
	return n + variable3
}
func main() {
	var n int
	fmt.Println("Выполнение странного задания")
	fmt.Println("Введите число")
	fmt.Scan(&n)
	fmt.Println(someFunc1(n), someFunc2(n), someFunc3(n))
	fmt.Println(someFunc3(someFunc2(someFunc1(n))))
}
