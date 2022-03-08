package main

import "fmt"

func main() {
	var currentYear = 2022
	if age := currentYear - 1998; age < 16 {
		fmt.Println("Kamu belum bisa membuat kartu sim")
	} else {
		fmt.Println("Kamu bisa membuat kartu sim")
	}
}
