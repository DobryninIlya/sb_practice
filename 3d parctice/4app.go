package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println(" Наполнение переменных особым образом")
	for {
		if a != 10 {
			a++
			continue
		}
		if b != 100 {
			b++
			continue
		}
		if c != 1000 {
			c++
			continue
		}
		break
	}
	fmt.Println(a, b, c)
}
