package main

import (
	"fmt"
	"strconv"
	"sync"
)

func input(wg *sync.WaitGroup) <-chan int {
	var inp string
	out := make(chan int)
	fmt.Scan(&inp)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for inp != "exit" {
			num, ok := strconv.Atoi(inp)
			if ok != nil {
				fmt.Println("Ошибка ввода. Ввод завершен!")
				break
			}
			fmt.Println("Ввод: ", inp)
			out <- num
			fmt.Scan(&inp)
		}
		close(out)
	}()
	return out
}

func square(ch <-chan int, wg *sync.WaitGroup) <-chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			num, ok := <-ch
			if !ok {
				break
			}
			num *= num
			out <- num
			fmt.Println("Квадрат: ", num)
		}
		close(out)
	}()
	return out
}

func multiply(ch <-chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			num, ok := <-ch
			if !ok {
				break
			}
			num *= 2
			fmt.Println("Произведение: ", num)
		}

	}()
}

func main() {
	var wg sync.WaitGroup
	ch1 := input(&wg)       // Добавление
	ch2 := square(ch1, &wg) // Возведение в квадрат
	multiply(ch2, &wg)
	wg.Wait()

}
