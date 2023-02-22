package main

import "fmt"

func main() {
	var moneyCountInput int
	var twoLastDigits int
	fmt.Println("Банкомат")
	fmt.Println("Введите сумму снятия со счета:")
	fmt.Scan(&moneyCountInput)
	if moneyCountInput > 100000 {
		fmt.Println("Введенная сумма превышает порог для снятия наличных")
		return
	}
	twoLastDigits = moneyCountInput % 100
	if twoLastDigits != 0 {
		fmt.Println("Банкомат выдает купюры только номиналом 100 рублей.\nВведенная вами сумма не может быть выдана")
		return
	} else {
		fmt.Print("Купюры выданы")
	}

}
