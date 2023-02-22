package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func printNums(quit chan os.Signal, wg *sync.WaitGroup) {
	x := 0
	for {
		x += 1
		select {
		case <-quit:
			fmt.Println("Прерывание. Выхожу из программы...")
			wg.Done()
			return
		default:
			fmt.Println(x, "|", x*x)
			time.Sleep(time.Second / 2)
		}
	}
}

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	var wg sync.WaitGroup
	wg.Add(1)
	go printNums(quit, &wg)
	wg.Wait()

}
