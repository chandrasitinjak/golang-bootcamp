package main

import (
	"fmt"
	"sync"
)

func main() {
	fruits := []string{"apple", "mangga", "durian", "rambutan"}

	var wg sync.WaitGroup
	for index, value := range fruits {
		wg.Add(1)
		go printFruit(index, value, &wg)
	}

	wg.Wait()
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("Index => %d, fruit => %s\n", index, fruit)
	wg.Done()
}
