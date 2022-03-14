package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string, grade int) {
	s.Name = name
	s.Grade = grade
}

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel : ", reflectValue.Int())
	}

	//access value with interface
	fmt.Println("nilai variabel : ", reflectValue.Interface().(int))

	var s1 = &student{
		Name:  "John wick",
		Grade: 2,
	}

	var reflectValue2 = reflect.ValueOf(s1)

	var method = reflectValue2.MethodByName("SetName")

	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
		reflect.ValueOf(23),
	})
	fmt.Println(reflectValue2)

	//fmt.Println("nama : ", s1.Name)
}
