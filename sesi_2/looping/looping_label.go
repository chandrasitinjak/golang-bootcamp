package main

import "fmt"

func main() {
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("perulangan i ke :", i)

		for j := 0; j <= 3; j++ {
			fmt.Println("data : ", j)
			if i == 2 {
				break outerLoop
			}
		}

		fmt.Println()
	}
}
