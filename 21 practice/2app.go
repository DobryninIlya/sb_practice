package main

import (
	"fmt"
	"log"
)

func decor(A func(int, int)) {
	defer A(100, 300)
	fmt.Println("Some pre-decorator")
}

func main() {
	decor(func(i int, i2 int) { fmt.Println(i + i2) })
	decor(func(i int, i2 int) { log.Println("Error code ", i2) })
	decor(func(i int, i2 int) { fmt.Println("Exiting") })
}
