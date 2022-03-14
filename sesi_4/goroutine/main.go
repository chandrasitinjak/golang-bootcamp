package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//go goroutine()
	var wg sync.WaitGroup
	fmt.Println("Main execution started")
	wg.Add(1)
	go firstProcess(8, &wg)
	secondProcess(8)

	fmt.Println("No of Goroutine:", runtime.NumGoroutine())

	fmt.Println("Main execution ended")
	wg.Wait()
}

func goroutine() {
	fmt.Println("Hello")

}

func firstProcess(index int, wg *sync.WaitGroup) {

	fmt.Println("first process func started")
	for i := 1; i <= index; i++ {

		fmt.Println("i=", i)

	}

	wg.Done()

	fmt.Println("First process func ended")
}

func secondProcess(index int) {
	fmt.Println("second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}

	fmt.Println("Second process func ended")
}
