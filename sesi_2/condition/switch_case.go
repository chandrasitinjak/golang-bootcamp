package main

import "fmt"

func main() {
	var score = 8

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	case 4:
		fmt.Println("not bad")
	default:
		{
			fmt.Println("default")
			fmt.Println("default 2")
		}

	}

	//fallthrough keyword
	data := 4

	switch {
	case data < 4:
		fmt.Println("kurang")
	case data >= 4 && data <= 10:

		fmt.Println("lebih dari 4 kurang dari 10")
		fallthrough

	default:
		{
			fmt.Println("mantap")
			fmt.Println("mantap2 ")
		}
	}

	//Nested condition
	var goal = 0
	if goal > 7 {
		switch {
		case goal < 10:
			fmt.Println("keren bana")
		default:
			fmt.Println("keren bini")
		}
	} else {
		if goal == 5 {
			fmt.Println("nice try")
		} else {
			if goal == 0 {
				fmt.Println("just go home")
			}
		}
	}

}
