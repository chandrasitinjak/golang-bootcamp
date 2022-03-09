package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee Employee

	employee.name = "Dedi"
	employee.age = 123
	employee.division = "Developer"

	fmt.Println(employee)
	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)

	/*
		Initializing struct
	*/

	var employee1 = Employee{}
	employee1.name = "ucok"
	employee1.age = 23
	employee1.division = "deve"

	var employee2 = Employee{
		name:     "chandra",
		age:      233,
		division: "olahraga",
	}

	fmt.Println(employee1)
	fmt.Println(employee2)

	/*
		Pointer to a struct
	*/

	var employee3 = Employee{
		name:     "employee3",
		age:      3,
		division: "emp3",
	}

	var employee4 *Employee = &employee3

	fmt.Println(employee3.name)
	fmt.Println(employee4.name)

	fmt.Println(strings.Repeat("#", 30))

	employee4.name = "ucok"
	fmt.Println(employee3.name)
	fmt.Println(employee4.name)

	/*
		Embedded struct
	*/
	
}
