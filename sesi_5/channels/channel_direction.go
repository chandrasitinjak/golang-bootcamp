package main

import "fmt"

func main() {
	c := make(chan string)

	students := []string{"Ariel", "Mailo", "Indah"}

	for _, v := range students {
		go intro(v, c)
	}

	for i := 1; i <= 3; i++ {
		prints(c)
	}
	close(c)
}

func prints(c <-chan string) {
	fmt.Println(<-c)
}

func intro(students string, c chan<- string) {
	result := fmt.Sprintf("Hai, my name is %s", students)
	c <- result
}
