package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString_1 = `
	{
	"full_name": "dedi chandra",
	"email": "dedi@gmail.com",
	"age": 23}`

	var result map[string]interface{}

	var err = json.Unmarshal([]byte(jsonString_1), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
	fmt.Println("full name : ", result["full_name"])
	fmt.Println("email : ", result["email"])
	fmt.Println("age: ", result["age"])

}