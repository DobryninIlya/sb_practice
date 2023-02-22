package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, D float64
	fmt.Println("Решение квадратного уравнения")
	fmt.Println("Введите коэффициенты a, b, c")
	fmt.Scan(&a, &b, &c)
	D = math.Pow(b, 2) - 4*a*c
	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / 2 * a
		x2 := (-b - math.Sqrt(D)) / 2 * a
		fmt.Println("Корни квадратного уравнения равны: ", x1, x2)
	} else if D == 0 {
		x1 := -b / 2 * a
		fmt.Println("Корень квадратного уравнения:", x1)
	} else {
		fmt.Println("Уравнение не имеет действительных корней!")
	}

}
