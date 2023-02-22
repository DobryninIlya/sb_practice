package main

import (
	"fmt"
	"math/rand"
	"time"
)

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int()
	fmt.Println(randNum, isEven(randNum))
}
