package main

import "fmt"

func main() {
	//variadic function is a function that can have unlimited input parameter, but only in last parameter
	studentList := print("dedi", "chandra", "sitinjak", "marga")
	fmt.Printf("%v\n", studentList)

	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	data := sum(numberList...)
	fmt.Println(data)
}

func print(name ...string) []map[string]string {
	var result []map[string]string

	for i, v := range name {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}

		result = append(result, temp)
	}

	return result
}

func sum(number ...int) int {
	total := 0

	for _, v := range number {
		total += v
	}

	return total
}
