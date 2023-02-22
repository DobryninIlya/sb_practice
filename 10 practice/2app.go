package main

import "fmt"

func main() {
	var summ float64
	var persent, years int
	fmt.Println("Введите сумму вклада, ежемесячный процент капитализации и количество лет")
	fmt.Scan(&summ, &persent, &years)
	bankSumm := 0.0
	for i := 0; i < years*12; i++ {
		summ += summ / 100.00 * float64(persent)
		temp := summ
		summ = float64(int(summ*100.00)) / 100.0
		bankSumm += temp - summ
	}
	fmt.Println(summ, bankSumm)
}
