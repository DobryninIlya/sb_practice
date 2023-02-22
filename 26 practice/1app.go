package main

import (
	"fmt"
	"io"
	"os"
)

func writeFile(text string, filename string) {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)
}

func readFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		_, err := file.Read(data)
		if err == io.EOF {
			break
		}
	}
	return string(data)
}

func main() {
	args := os.Args[1:]
	if len(args) > 3 {
		fmt.Println("Ошибка аргументов. \nПример: first.txt second.txt result.txt")
	} else if len(args) == 2 {
		writeFile("some first string", args[0])
		writeFile("last string in result file", args[1])
		string1 := readFile(args[0])
		string2 := readFile(args[1])
		fmt.Println(string1 + "\n" + string2)
	} else if len(args) == 1 {
		string1 := readFile(args[0])
		fmt.Println(string1)
	} else {
		writeFile("some first string", args[0])
		writeFile("last string in result file", args[1])
		string1 := readFile(args[0])
		string2 := readFile(args[1])
		writeFile(string1+"\n"+string2, args[2])

	}
}
