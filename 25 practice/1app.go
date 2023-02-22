package main

import (
	"flag"
	"fmt"
)

func main() {
	var fullStr string
	var subStr string
	flag.StringVar(&fullStr, "str", "", "set fullStr")
	flag.StringVar(&subStr, "substr", "", "set subStr")
	flag.Parse()
	k := 0
	for _, val := range fullStr {
		if val == rune(subStr[k]) {
			k++
			if len(subStr) == k {
				break
			}

		} else {
			k = 0
		}
	}
	if k != 0 {
		fmt.Println("Содержит")
	} else {
		fmt.Println("Не содержит")
	}
}
