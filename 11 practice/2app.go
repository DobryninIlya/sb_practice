package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "a10 10 20b 20 30c30 30 dd"
	firstIncome := true
	fmt.Println("Строки в десятичном формате:")
	for strings.Contains(str, " ") {
		spaceIndex := strings.Index(str, " ")
		word := str[:spaceIndex]
		str = str[spaceIndex+1:]

		if !strings.Contains(str, " ") && firstIncome {
			str = str + " "
			firstIncome = false
		}
		_, error := strconv.Atoi(word)
		if error == nil {
			fmt.Print(word + " ")
		}

	}

}
