package main

import "fmt"

func main() {
	var flag bool
	fmt.Println("Угадай число. Акинатор")
	fmt.Println("Введенное число больше или равно 5 ?")
	fmt.Scan(&flag)
	if flag {
		fmt.Println("Введенное число больше или равно 7?")
		fmt.Scan(&flag)
		if flag {
			fmt.Println("Введенное число больше или равно 9 ?")
			fmt.Scan(&flag)
			if flag {
				fmt.Println("Введенное число 10?")
				fmt.Scan(&flag)
				if flag {
					fmt.Println("Ха! Я угадал! Ответ 10")
					return
				} else {
					fmt.Println("Ответ 9!")
					return
				}
			}
			fmt.Println("Введенное число 8?")
			fmt.Scan(&flag)
			if flag {
				fmt.Println("Ответ 8")
				return
			} else {
				fmt.Println("Ответ 7")
				return
			}
		}
		fmt.Println("Введенное число 6?")
		fmt.Scan(&flag)
		if flag {
			fmt.Println("Ответ 6")
			return
		} else {
			fmt.Println("Ответ 5")
			return
		}
	} else {
		fmt.Println("Введенное число меньше или равно 3?")
		fmt.Scan(&flag)
		if flag {
			fmt.Println("Введенное число меньше или равно 2 ?")
			fmt.Scan(&flag)
			if flag {
				fmt.Println("Введенное число 1?")
				if flag {
					fmt.Println("Ответ 1")
					return
				} else {
					fmt.Println("Ответ 2")
					return
				}
			} else {
				fmt.Println("Введенное число  3!")
			}
		} else {
			fmt.Println("Введенное число  4!")
		}

	}

}
