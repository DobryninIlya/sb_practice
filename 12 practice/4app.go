// 1ая задача итак использован ioutil
package main

import (
	//"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	//var b bytes.Buffer
	filename := "output.txt"
	str, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка открытия файла ", err)
	}
	fmt.Printf("%s\n", str)
}
