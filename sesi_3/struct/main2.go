package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employeee struct {
	division string
	person   Person
}

func main() {
	//var employee Employee
	//var employee = Employee{}
	var employee = Employeee{
		division: "olahraga",
		person: Person{
			name: "chandra",
			age:  39,
		},
	}

	fmt.Println(employee)
	fmt.Printf("%+v\n", employee)

	/*
		Anonymous struct
		without full the field
	*/
	var employee1 = struct {
		person   Person
		division string
	}{}

	employee1.person = Person{
		name: "a",
		age:  23,
	}
	employee1.division = "mancap"

	fmt.Println(employee1)

	/*
		Anonymous struct
		with full the field
	*/

	var employee2 = struct {
		person   Person
		division string
	}{
		person: Person{
			name: "b",
			age:  231,
		},
		division: "kitakita",
	}

	fmt.Println(employee2)
}
