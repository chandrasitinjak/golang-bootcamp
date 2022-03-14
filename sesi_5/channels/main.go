package main

import "fmt"

func main() {
	c := make(chan string)

	go introduce("a", c)
	go introduce("b", c)
	go introduce("c", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	close(c)

}

func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result
}
