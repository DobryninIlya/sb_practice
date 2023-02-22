package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	var b bytes.Buffer
	b.WriteString("TOP SECRET INFO")
	filename := "secret.txt"
	if err := ioutil.WriteFile(filename, b.Bytes(), 0444); err != nil {
		panic(err)
	}
	f, err := os.Open("secret.txt")
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
