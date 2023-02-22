package main

import "fmt"

func main() {
	var arr [10]int
	chetn := 0
	for i, _ := range arr {
		fmt.Println("Введите переменную ", i+1)
		fmt.Scan(&arr[i])
	}
	for _, val := range arr {
		if val%2 == 0 {
			chetn++
		}
	}
	fmt.Printf("Четных %v,\n Нечетных %v", chetn, 10-chetn)
}
