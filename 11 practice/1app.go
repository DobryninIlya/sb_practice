package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go is an Open source programming Language that makes it Easy \nto build simple, reliable, and efficient Software"
	firstIncome := true
	count := 0
	for strings.Contains(str, " ") {
		spaceIndex := strings.Index(str, " ")
		word := str[:spaceIndex]
		str = str[spaceIndex+1:]

		if !strings.Contains(str, " ") && firstIncome {
			str = str + " "
			firstIncome = false
		}
		if word == strings.Title(word) {
			count++
		}
	}
	fmt.Printf("Слов с большой буквы: %v", count)
}
