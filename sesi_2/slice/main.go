package main

import "fmt"

func main() {
	//declare a slice

	var fruits = []string{"apple", "manggo", "banana"}
	_ = fruits

	var fruit = make([]string, 3)
	_ = fruit

	var friuits1 = make([]string, 3)
	friuits1 = append(friuits1, "apple", "mangga", "nanas")
	fmt.Printf("%#v\n", friuits1)

	var fruits2 = []string{"asam", "jeruk", "pisang"}
	//var fruits3 = []string{"asam1", "jeruk1", "pisang"}
	fruits2 = append(fruits2, friuits1...)
	fmt.Printf("%#v\n", fruits2)

	//combininig slice and append
	var buah = []string{"mangga", "apel", "nanas", "pisang", "alpukat", "rambutan"}
	buah = append(buah[:3], "sawit")

	fmt.Printf("%#v\n", buah)
}
