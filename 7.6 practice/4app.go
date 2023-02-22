package main

import (
	"fmt"
)

func main() {
	fmt.Println("Счастливые билеты")
	max := 0
	lastHappyTicket := 0
	for i := 100000; i < 1000000; i++ {
		rightSide := 0
		leftSide := 0
		rightSide += i / 100000 % 10
		rightSide += i / 10000 % 10
		rightSide += i / 1000 % 10
		leftSide += i / 100 % 10
		leftSide += i / 10 % 10
		leftSide += i % 10
		if leftSide == rightSide {
			if i-lastHappyTicket >= max && lastHappyTicket != 0 {
				max = i - lastHappyTicket
			}
			lastHappyTicket = i
		}
	}
	fmt.Println("Минимальное количество билетов, которое нужно кпить, чтобы среди них оказаля счастливый:", max)
}
