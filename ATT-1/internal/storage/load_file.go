package storage

import (
	"io"
	"log"
	"os"
	"strings"
)

func LoadFile(path string) {
	fileString, err := readFile(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	rows := strings.Split(fileString, "/n")
	for i, row := range rows {
		params := strings.Split(row, ",")
		if len(params) != 6 {
			log.Fatalf("Входной файл имеет неподдерживаемый формат около строки %v", i)
			return
		}
	}
}

func readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data := make([]byte, 0, 64)

	for {
		_, err := file.Read(data)
		if err == io.EOF {
			break
		}
	}
	return string(data), nil

}
