package main

import (
	"fmt"
)

func main() {
	employeeSalary := map[string]int{
		"Ram":     12000,
		"Hari":    15000,
		"Krishna": 90000,
	}
	employee := "Hari"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)
}
