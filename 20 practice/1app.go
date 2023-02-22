package main

import "fmt"

const (
	rows = 3
	cols = rows
)

func getDeterminer(arr [rows][cols]int) int {
	var a [rows][cols + cols - 1]int
	var result int
	for i := 0; i < len(a); i++ { // По правилу Саррюса дублируем первые столбцы в новую матрицу для удобства
		for j := 0; j < cols; j++ {
			a[i][j] = arr[i][j]
		}
		for j := 0; j < cols-1; j++ { // По правилу Саррюса дублируем первые столбцы
			a[i][j+3] = a[i][j]
		}
	}
	margin := 0
	for _, val := range a {
		fmt.Println(val)
	}
	for i := 0; i < cols+cols-1; i++ {
		if i < 3 {
			result += a[i-margin][i] * a[i+1-margin][i+1] * a[i+2-margin][i+2] // Положительные слагаемые
		}
		if i >= 2 {
			result -= a[i-margin][i] * a[i-margin+1][i-1] * a[i-margin+2][i-2] // Отрицательные слагаемые
		}
		margin++
	}
	return result
}

func main() {
	matrix := [rows][cols]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(getDeterminer(matrix))
}
