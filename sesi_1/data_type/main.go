package main

import "fmt"

func main() {
	var first uint8 = 89
	var second int8 = -12

	fmt.Printf("Tipe data first : %T\n", first)
	fmt.Printf("Tipe data second : %T\n", second)

	//float

	var decimalNumber float32 = 3.63

	fmt.Printf("decimal number %f\n", decimalNumber)
	fmt.Printf("decimal number %.3f\n", decimalNumber)

	//boolean
	data := true

	if data {
		fmt.Println("benar")
	} else {
		fmt.Println("salah")
	}

	/*
		zero data != nil
		zero data terpat pada string, boolean, numerik


		nil terdapat pada :
		pointer,tipe data function,slice, map,channel, interface
	*/

}
