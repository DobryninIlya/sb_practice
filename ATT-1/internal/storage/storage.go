package storage

import (
	city "main.go/internal/city"
	"sync"
)

type MainSotrage struct {
	Store map[int]*city.City
	sync.Mutex
}

func (ms *MainSotrage) Create() {

}

func (ms *MainSotrage) GetById() {

}