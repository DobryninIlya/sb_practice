package main

import (
	"fmt"
	"math"
)

func main() {
	var x, n int
	ex := math.E
	previous := 0.0
	i := 0
	fact := 1
	fmt.Println("Введите x, n")
	fmt.Scan(&x, &n)
	accuracy := 1 / math.Pow10(n)
	for {
		if i > 0 {
			fact *= i
			ex += math.E * math.Pow(float64(x-1), float64(i)) / float64(fact)
		}
		if math.Abs(ex-float64(previous)) < accuracy {
			fmt.Println(ex)
			fmt.Println(math.E)
			break
		}
		i++
		previous = ex
	}
}
