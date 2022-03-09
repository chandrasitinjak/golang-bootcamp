package main

import "belajar-go/helper"

func main() {
	helper.Greet()
	//helper.Person.Invokegreet()
	var person = helper.Person{}
	person.Invokegreet()
}
