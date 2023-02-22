package main

import (
	"fmt"
	"math"
)

func main() {
	overflow8 := 0
	overflow16 := 0
	for i := 0; i <= math.MaxUint32; i++ {
		if i > math.MaxUint8 {
			overflow8++
		}
		if i > math.MaxUint16 {
			overflow16++
		}
	}
	fmt.Println(overflow8, overflow16)
}
