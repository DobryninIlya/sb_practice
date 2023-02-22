package main

import "fmt"

func main() {
	var bill5, bill10, bill20, input int
	for {
		fmt.Println("Введите сумму")
		fmt.Scan(&input)
		switch input {
		case 5:
			bill5++
			fmt.Println("Сдача есть. Торговать можно")
		case 10:
			bill10++
			if bill5 > 0 {
				fmt.Println("Сдача есть. Торговать можно")
				bill5--
			} else {
				fmt.Println("Сдачи нет.")
				return
			}
		case 20:
			bill20++
			if bill5 > 0 && bill10 > 0 {
				fmt.Println("Сдача есть. Торговать можно")
				bill5--
				bill10--
			} else if bill5 >= 3 {
				fmt.Println("Сдача есть. Торговать можно")
				bill5 -= 3
			} else {
				fmt.Println("Сдачи нет. Торговля невозможна")
				return
			}
		default:
			return
		}

	}
}
