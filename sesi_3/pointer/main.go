package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstNumber int = 10
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber value : ", firstNumber)
	fmt.Println("firstNumber address : ", &firstNumber)

	//menyimpan nilai dari alamat yang disimpan
	fmt.Println("secondNumber value : ", *secondNumber)
	//menyimpan alamat dari firstNumber
	fmt.Println("secondNumber address : ", secondNumber)
	//alamat dari secondNumber
	fmt.Println("secondNumber address ?  : ", &secondNumber)

	/*
		Changes value through pointer
	*/

	var firstPerson = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("nilai firstPerson ", firstPerson)
	fmt.Println("alamat address firstPerson ", &firstPerson)

	fmt.Println("nilai dari alamat secondNumber ", *secondPerson)
	fmt.Println("nilai secondPerson ", secondPerson)
	fmt.Println("alamat secondPerson ", &secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Sitinjak"

	fmt.Println("nilai firstPerson ", firstPerson)
	fmt.Println("alamat address firstPerson ", &firstPerson)

	fmt.Println("nilai dari alamat secondNumber ", *secondPerson)
	fmt.Println("nilai secondPerson ", secondPerson)
	fmt.Println("alamat secondPerson ", &secondPerson)

	/*
		Pointer as param
	*/
	var a = 10
	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("Before:", a)

}

func changeValue(number *int) {
	*number = *number * 2
}
