package main

import "fmt"

func main() {
	var width, height int
	fmt.Println("Шахмотная доска")
	fmt.Println("Введите ширину и длину")
	fmt.Scan(&width, &height)
	margin := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if (j+margin)%2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		margin++
		fmt.Print("\n")

	}

}
