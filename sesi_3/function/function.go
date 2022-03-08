package main

import (
	"fmt"
	"math"
)

func main() {
	greet("Dedi", 22)
	data := withReturn(21, 23)
	fmt.Println(data)

	var area, circumference = calculate(15)
	fmt.Println(area, circumference)
	var area_, circumference_ = calculate2(125)
	fmt.Println(area_, circumference_)
}

func greet(name string, age int) {
	fmt.Printf("Hello there! My name is %s and I am %d years old\n", name, age)
}

func withReturn(number1, number2 int) int {
	return number1 + number2
}

//function return multiple values
func calculate(d float64) (area float64, circumFerence float64) {
	area = math.Pi * math.Pow(d/2, 2)

	circumFerence = math.Pi * d

	return
}

//function return multiple values and predefined return value
func calculate2(d float64) (float64, float64) {
	var area float64 = math.Pi * math.Pow(d/2, 2)

	var circumFerence = math.Pi * d

	return area, circumFerence
}
