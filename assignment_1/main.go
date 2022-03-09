package main

import (
	"assignment_1/helper"
	"os"
	"strconv"
)

func main() {
	id_user, _ := strconv.Atoi(os.Args[1])

	edpKoinwork := []helper.EdpKoinWork{
		{
			Employee: helper.Employee{
				Person: helper.Person{
					Name: "alvin",
				},
				Address:    "balige",
				Profession: "Backend Engineer",
			},
			Motivation: "Go is one of the most famous programming language now",
		},
		{
			Employee: helper.Employee{
				Person: helper.Person{
					Name: "evrin",
				},
				Address:    "jakarta",
				Profession: "Backend Engineer",
			},
			Motivation: "Go is my favorite language",
		},
		{
			Employee: helper.Employee{
				Person: helper.Person{
					Name: "khairul",
				},
				Address:    "jakarta",
				Profession: "Backend Engineer",
			},
			Motivation: "i love go",
		},
		{
			Employee: helper.Employee{
				Person: helper.Person{
					Name: "faikar",
				},
				Address:    "jakarta",
				Profession: "Backend Engineer",
			},
			Motivation: "Go is interesting",
		},
		{
			Employee: helper.Employee{
				Person: helper.Person{
					Name: "mghifari",
				},
				Address:    "jakarta",
				Profession: "Backend Engineer",
			},
			Motivation: "i love new technology",
		},
		{
			Employee: helper.Employee{
				Person: helper.Person{
					Name: "dedi",
				},
				Address:    "riau",
				Profession: "Backend Engineer",
			},
			Motivation: "Go is simple, clean, and fast",
		},
	}

	//fmt.Println(id_user)

	helper.FindBiodataUser(edpKoinwork, id_user)

}
