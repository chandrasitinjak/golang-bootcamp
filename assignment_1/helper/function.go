package helper

import "fmt"

func FindBiodataUser(edpKoinwork []EdpKoinWork, id_user int) {
	for index, value := range edpKoinwork {
		var temp int = index
		temp += 1
		if temp == id_user {
			fmt.Println("name : ", value.Employee.Person.Name)
			fmt.Println("address : ", value.Employee.Address)
			fmt.Println("profession : ", value.Employee.Profession)
			fmt.Println("motivation : ", value.Motivation)
		}
	}
}
