package main

import "fmt"

func main() {
	callDeferFunc()

	fmt.Println("hei yoo everyone")
}

func callDeferFunc() {
	defer deferFunc()
	fmt.Println("hello world")
}

func deferFunc() {
	fmt.Println("defer function starts to execute")
}
