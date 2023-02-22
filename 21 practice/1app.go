package main

import "fmt"

func function(x int16, y uint8, z float32) float32 {
	return 2*float32(x) + float32(y*y) - 3/z
}

func main() {
	fmt.Println(function(10, 5, 1.27))
}
