package main

import "fmt"

func main() {
	//create alias
	type number = int

	var data number = 13

	fmt.Println(data)
}
