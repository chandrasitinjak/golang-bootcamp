package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// sayHello()
	//comentar one line

	/*
		comentar multiple line
	*/

	//variabel with data type
	var name string = "Dedi Chandra"
	var age int = 123

	var name_ = "Dedi Chandra"
	var age_ = 1231

	fmt.Println("nama ", name)
	fmt.Println("age ", age)

	fmt.Println("nama ", name_)
	fmt.Println("age ", age_)

	//short declarations

	satu, dua, tiga := 1, 2, 3
	empat, lima, enam := 4, 5, 6

	fmt.Println(satu, dua, tiga, empat, lima, enam)

}
