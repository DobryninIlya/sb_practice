package main

import (
	"fmt"
)

func main() {
	var firstExamScore int
	var secondExamScore int
	var thirdExamScore int
	enterScore := 275
	fmt.Println("Баллы ЕГЭ")
	fmt.Println("Введите результат первого экзамена")
	fmt.Scan(&firstExamScore)
	fmt.Println("Введите результат второго экзамена")
	fmt.Scan(&secondExamScore)
	fmt.Println("Введите результат третьего экзамена")
	fmt.Scan(&thirdExamScore)
	fmt.Println("Сумма проходных баллов: ", enterScore)
	summa := firstExamScore + secondExamScore + thirdExamScore
	fmt.Println("Количество набранных баллов: ", summa)

	if summa >= enterScore {
		fmt.Println("Вы поступили!!!")
	} else {
		fmt.Println("Вы не поступили")
	}
}
