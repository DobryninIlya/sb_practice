package main

import "fmt"

func main() {
	var basket1Capacity, basket2Capacity, basket3Capacity, basket1, basket2, basket3 int
	fmt.Println("Корзины с яблоками")
	fmt.Println("Введите сколько яблок помещается в каждую корзину")
	fmt.Scan(&basket1Capacity, &basket2Capacity, &basket3Capacity)
	for {
		if basket1 != basket1Capacity {
			fmt.Println("Корзина 1 пополнилась на 1 яблоко")
			basket1++
			continue
		}
		if basket2 != basket2Capacity {
			fmt.Println("Корзина 2 пополнилась на 1 яблоко")
			basket2++
			continue
		}
		if basket3 != basket3Capacity {
			fmt.Println("Корзина 3 пополнилась на 1 яблоко")
			basket3++
			continue
		}
		break
	}

}
