package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("12 practice/output.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла ", err)
	}
	defer f.Close()
	fInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fInfo.Size())
	if _, err := io.ReadFull(f, buf); err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", buf)
}
