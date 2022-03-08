package main

import (
	"fmt"
	"strings"
)

func main() {
	//Closure is a anonymous function that save as a variable or as a return value from function
	//declare closure in variable
	var evenNUmbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	var numbers = []int{1, 4, 6, 2, 3, 5, 8, 12, 512, 2321}
	fmt.Println(evenNUmbers(numbers...))

	//Closure immediately-invoked function expressin (IIFE)
	var isPalindrome = func(str string) bool {

		for i := 0; i < len(str)/2; i++ {
			if string(byte(str[i])) != string(byte(str[len(str)-i-1])) {
				return false
			}
		}

		return true
	}("kataks")

	fmt.Println(isPalindrome)

	//closure as a return value
	var studentLists = []string{"ada", "adi", "adu", "ade", "ado"}

	var find = findStudent(studentLists)

	fmt.Println(find("adas"))

	var angka = []int{2, 5, 8, 10, 3, 99, 23}
	var cari = findOddNumbers(angka, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("total number odd", cari)
}

func findOddNumbers(nums []int, callback func(int) bool) int {
	var totalOddNumber = 0
	for _, val := range nums {
		if callback(val) {
			totalOddNumber++
		}
	}

	return totalOddNumber
}

func findStudent(studentss []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range studentss {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}

		if student == "" {
			return fmt.Sprintf("%s doesn't exists!!!\n", s)
		}

		return fmt.Sprintf("%s exists!!! at position %d\n", s, position)

	}
}
