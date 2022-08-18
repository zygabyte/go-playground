package main

import "fmt"

type Employee struct {
	Name string
	Id   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s - (%s)", e.Name, e.Id)
}

type Manager struct {
	Employee
	Reports []Employee
}

func main() {
	e := Employee{
		Name: "James White",
		Id:   "21890234",
	}

	fmt.Println(e.Description())
}
