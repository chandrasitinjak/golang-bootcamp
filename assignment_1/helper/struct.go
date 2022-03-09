package helper

type Person struct {
	Name string
}

type Employee struct {
	Person     Person
	Address    string
	Profession string
}

type EdpKoinWork struct {
	Employee   Employee
	Motivation string
}
