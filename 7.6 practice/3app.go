package main

import "fmt"

func main() {
	var count int
	fmt.Println("Елочка")
	fmt.Println("Введите высоту")
	fmt.Scan(&count)
	for i := 1; i <= count; i++ {
		spaces := count - i
		for j := spaces; j > 0; j-- {
			fmt.Print(" ")
		}
		for j := i; j < i+i*2-1; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
