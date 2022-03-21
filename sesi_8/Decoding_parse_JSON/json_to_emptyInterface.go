package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString_2 = `
	{
	"full_name": "dedi chandra",
	"email": "dedi@gmail.com",
	"age": 23}`

	var temp interface{}

	var err = json.Unmarshal([]byte(jsonString_2), &temp)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result = temp.(map[string]interface{})

	fmt.Println(result)
	fmt.Println("fullname : ", result["full_name"])
	fmt.Println("email : ", result["email"])
	fmt.Println("age : ", result["age"])
}
