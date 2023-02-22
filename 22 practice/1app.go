package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func main() {
	var arr [n]int
	var input int
	rand.Seed(time.Now().UnixNano())
	for i, _ := range arr {
		arr[i] = rand.Intn(n)
	}
	fmt.Println(arr)
	fmt.Scan(&input)
	flag := 0
	for i, val := range arr {
		if val == input {
			flag = n - i - 1
			break
		}
	}
	fmt.Println(flag)
}
