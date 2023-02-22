package main

import (
	"fmt"
	"math"
)

func main() {
	var digit1, digit2 int16
	unsigned := false
	fmt.Println("Введите 2 числа")
	fmt.Scan(&digit1, &digit2)
	multplication := int32(digit1) * int32(digit2)
	if multplication > 0 {
		unsigned = true
	} else {
		multplication = -multplication
	}
	switch {
	case multplication <= math.MaxUint8 && unsigned:
		fmt.Println("uint8")
	case multplication <= math.MaxUint16 && unsigned:
		fmt.Println("uint16")
	case uint32(multplication) <= math.MaxUint32 && unsigned:
		fmt.Println("uint32")
	case multplication <= math.MaxInt8 && !unsigned:
		fmt.Println("int8")
	case multplication <= math.MaxInt16 && !unsigned:
		fmt.Println("int16")
	case multplication <= math.MaxInt32 && !unsigned:
		fmt.Println("int32")
	}

}
