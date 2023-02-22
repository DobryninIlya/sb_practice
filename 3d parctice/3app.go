package main

import "fmt"

func main() {
	var price, discount float64
	fmt.Println("Расчет суммы скидки")
	fmt.Println("Введите цену товара и скидку:")
	fmt.Scan(&price, &discount)
	currentDiscount := 0.0
	discountPersent := discount / price * 100 // Скидка в процентах
	if discount >= 2000 || discountPersent >= 30 {
		discountPersent := price / 100 * 30
		if discountPersent >= 2000 {
			currentDiscount = 2000
		} else {
			currentDiscount = discountPersent
		}
	} else {
		currentDiscount = discount
	}
	fmt.Println(currentDiscount)
}
