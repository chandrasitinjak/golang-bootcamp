package main

import "fmt"

func main() {
	//declare a map
	var person map[string]string
	person = map[string]string{}

	person["name"] = "a"
	person["age"] = "23"
	person["address"] = "jakarta"

	fmt.Println("name", person["name"])
	fmt.Println("age", person["age"])
	fmt.Println("address", person["address"])

	var manusia = map[string]string{
		"name":    "ahc",
		"age":     "23",
		"address": "jauh",
	}

	fmt.Println("name", manusia["name"])
	fmt.Println("age", manusia["age"])
	fmt.Println("address", manusia["address"])

	for idx, value := range manusia {
		fmt.Println("id ", idx, "value", value)
	}

	delete(manusia, "age")

	for idx, value := range manusia {
		fmt.Println("id ", idx, "value", value)
	}

	value, exist := manusia["age"]
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("value has been delted")
	}

	//combining map and slice
	var data = []map[string]string{
		{"name": "lala", "age": "1"},
		{"name": "lili", "age": "2"},
		{"name": "lulu", "age": "3"},
		{"name": "lule", "age": "4"},
	}

	for i, value := range data {
		fmt.Println("index", i, "value", value)
	}
}
