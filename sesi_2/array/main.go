package main

import "fmt"

func main() {

	//declar array
	var number [4]int
	number = [4]int{1, 2, 3, 4}
	fmt.Printf("%#v\n", number)

	//multidimension array
	numbers := [2][4]int{{1, 2, 3, 4}, {6, 7, 8, 9}}

	for _, arr := range numbers {
		for _, value := range arr {
			fmt.Println(value)
		}

		fmt.Println("####")
	}
}
