package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	city := "jakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	}

	str1 := "makan"
	str2 := "mănca"

	//function len count a string by len of byte
	fmt.Println(len(str1), len(str2))

	//function utf8.RuneCountInString count a string by char
	fmt.Println(utf8.RuneCountInString(str1), utf8.RuneCountInString(str2))

	//looping a string
	for i, value := range str2 {
		fmt.Println("index", i, "value", string(value))
	}
}
