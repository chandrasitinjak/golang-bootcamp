package main

import "fmt"

func main() {
	c := make(chan string)

	students := []string{"arieal", "mailo", "indah"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Students", student)
			result := fmt.Sprintf("Hai my name is %s\n", student)
			c <- result
		}(v)
	}

	for i := 1; i <= 3; i++ {
		fmt.Println(<-c)
		//print(c)
	}

	/*
		c chan string => send and receive data
		c chan <- string => paramter c only send data
		c <- chan string => parameter c only receive data
	*/
}

func print(c chan string) {
	fmt.Println(<-c)
}
