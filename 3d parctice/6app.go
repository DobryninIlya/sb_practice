package main

import (
	"fmt"
)

func main() {
	fmt.Println("Лифт")
	notEnd := true
	flag10 := true
	flag7 := true
	flag4 := true
	for notEnd {
		currFlat := 0
		fmt.Println("Лифт на 1 этаже и начинает ехать вверх!")
		for currFlat < 24 {
			currFlat++
		}
		if currFlat == 24 {
			fmt.Println("Лифт приехал на последний этаж и начинает движение вниз")
			cabine := 0
			for ; currFlat >= 0; currFlat-- {
				if cabine == 2 {
					continue
				}
				if currFlat == 10 && flag10 {
					cabine++
					flag10 = false
					fmt.Println("В лифт вошел человек с 10 этажа")
				}

				if currFlat == 7 && flag7 {
					cabine++
					flag7 = false
					fmt.Println("В лифт вошел человек с 7 этажа")
				}

				if currFlat == 4 && flag4 {
					cabine++
					flag4 = false
					fmt.Println("В лифт вошел человек с 4 этажа")
					notEnd = false
				}

				if currFlat == 0 {
					fmt.Println("Лифт на 1 этаже")
				}
			}
		}

	}

}
