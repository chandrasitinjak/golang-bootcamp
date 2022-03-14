package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("invoke with dever")
	fmt.Println("before exiting")
	//os.Exit(1)
}
