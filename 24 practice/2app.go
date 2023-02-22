package main

import (
	"fmt"
	"strings"
)

func parseTest(sentences [4]string, chars [5]rune) [][]int {
	result := make([][]int, 4, 4)
	for i, sentence := range sentences {
		words := strings.Fields(sentence)
		lastWord := strings.ToUpper(words[len(words)-1])
		for _, char := range chars {
			result[i] = append(result[i], strings.Index(lastWord, string(char)))
			fmt.Println(strings.Index(lastWord, string(char)))
		}
	}

	return result
}

func main() {
	sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := [5]rune{'H', 'E', 'L', 'П', 'М'}
	result := parseTest(sentences, chars)
	for _, arr := range result {
		fmt.Println(arr)
	}
}
