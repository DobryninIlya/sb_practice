package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Зеркальные билеты")
	result := 0
	for digit := 100000; digit < 1000000; digit++ {
		reversed := 0
		i := digit
		j := 0
		for i > 0 {
			rightDigit := i % 10
			i /= 10
			b := rightDigit * int(math.Pow(10, float64((5-j))))
			reversed += b
			j++
		}
		if reversed == digit {
			result++
		}
	}
	fmt.Println(result)
}
