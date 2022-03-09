package main

import "fmt"

type Persons struct {
	name string
	age  int
}

func main() {
	var people = []Persons{
		{name: "chandra1", age: 1},
		{name: "chandra2", age: 2},
		{name: "chandra3", age: 3},
	}

	for _, v := range people {
		fmt.Println(v)
	}

	/*
		Anonymous slice of struct
	*/
	var employee = []struct {
		person   Persons
		division string
	}{
		{
			person: Persons{
				name: "a",
				age:  1,
			},
			division: "aaaa"},

		{
			person: Persons{
				name: "mantapp",
				age:  232,
			},
			division: "salah",
		},
	}

	for _, v := range employee {
		fmt.Println(v)
	}
}
