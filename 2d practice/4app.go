package main

import "fmt"

func main() {
	var price, value1, value2, value3 int
	fmt.Println("Сумма без сдачи")
	fmt.Println("Введите стоимость товара и номиналы трех монет")
	fmt.Scan(&price, &value1, &value2, &value3)
	if price == value1 || price == value2 || price == value3 {
		fmt.Println("Оплата может быть произведена только лишь одной монетой")
	} else if price == value1+value2 {
		fmt.Println("Оплата может быть произведена только лишь монетами номиналом", value1, value2)
	} else if price == value2+value3 {
		fmt.Println("Оплата может быть произведена только лишь монетами номиналом", value2, value3)
	} else if price == value1+value3 {
		fmt.Println("Оплата может быть произведена только лишь монетами номиналом", value1, value3)
	} else if price == value1+value2+value3 {
		fmt.Println("Оплата может быть произведена только всеми тремя монетами!")
	} else {
		fmt.Println("Трех монет не достаточно для оплаты товара!")
	}
}
