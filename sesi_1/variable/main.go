package main

import "fmt"

func main() {
	/*
		declare with initialization
			var name string = "Airell"
			var age int = 23 /*
	*/

	/* declare without initialization
	var name string
	var age int

	name = "chandra"
	age = 25
	*/

	//declare without data type
	//var name = "chandra"
	//var age = 23

	//name := "chandra"
	//age := 23

	//multiple declaration variable
	var student1, student2, student3 string = "satu", "dua", "tiga"

	var name, age, clan = "chandra", 12, "sitinjak"

	fmt.Println(name, age, clan)
	//fmt.Println("Ini adalah nama ==>", name)
	//fmt.Println("Ini adalah age ==>", age)
	fmt.Println("Ini adalah data ==>", student1, student2, student3)
}
