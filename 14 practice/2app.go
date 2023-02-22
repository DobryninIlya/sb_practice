package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomPoint(n int) (int, int) {
	return rand.Intn(n), rand.Intn(n)
}

func convertPoint(x, y int) (int, int) {
	return 2*x + 10, -3*y + 5
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(convertPoint(getRandomPoint(10)))
	fmt.Println(convertPoint(getRandomPoint(10)))
	fmt.Println(convertPoint(getRandomPoint(10)))
}
