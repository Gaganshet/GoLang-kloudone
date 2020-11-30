package main

import "fmt"

type Employee struct {
	name   string
	id     int
	course string
	year   int
	salary float32
	grade  string
}

type Company struct {
	name        string
	address     string
	Established int
	details     Employee
}

func main() {
	result := Company{
		name:        "Kloudone",
		address:     "Chennai",
		Established: 2018,
		details:     Employee{"GAGAN", 200, "Golang", 2020, 30000, "A"},
	}
	fmt.Println("-----Details of the COMPANY------")
	fmt.Println("Name: ", result.name)
	fmt.Println("Address:", result.address)
	fmt.Println("Established:", result.Established)

	fmt.Println("\n----Details of Employee----")
	fmt.Println("Employee name--->", result.details.name)
	fmt.Println("ID---> ", result.details.id)
	fmt.Println("Course---> ", result.details.course)
	fmt.Println("Year---> ", result.details.year)
	fmt.Println("Salary---> ", result.details.salary)
	fmt.Println("Grade--->", result.details.grade)
}
