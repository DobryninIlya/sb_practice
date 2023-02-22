package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	var b bytes.Buffer
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		panic(err)
	}
	for in.Text() != "exit" {

		prefix := time.Now().Format(time.RFC3339) + " "
		b.WriteString(prefix + in.Text() + "\n")

		in.Scan()
		if err := in.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
		}
	}
	filename := "output.txt"
	if err := ioutil.WriteFile(filename, b.Bytes(), 0666); err != nil {
		panic(err)
	}

}
