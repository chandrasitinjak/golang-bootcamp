package main

import (
	"encoding/json"
	"fmt"
)

type Emp struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString_3 = `[
	{
	"full_name": "dedi chandra",
	"email": "dedi@gmail.com",
	"age": 23},
	{
	"full_name": "chahyadi",
	"email": "cahyadi@gmail.com",
	"age": 34}
	]`

	var result []Emp

	var err = json.Unmarshal([]byte(jsonString_3), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}
}
