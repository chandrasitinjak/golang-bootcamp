package main

import (
	"fmt"
)

func main() {
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"a", "b"}

	var v interface{}

	v = 20
	fmt.Printf("%T\n", v)

	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	fmt.Println(v)

	/*
		empty interface with map & slice
	*/
	rs := []interface{}{1, "a", true, 2, "b", true}

	rm := map[string]interface{}{
		"Name":   "a",
		"status": true,
		"age":    23,
	}

	_, _ = rs, rm

}
