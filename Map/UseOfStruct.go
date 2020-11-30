package main

import (
	"fmt"
)

type employee struct {
	salary  int
	country string
}

func main() {
	emp1 := employee{
		salary:  12000,
		country: "India",
	}
	emp2 := employee{
		salary:  14000,
		country: "India",
	}
	emp3 := employee{
		salary:  13000,
		country: "India",
	}
	employeeInfo := map[string]employee{
		"Ganesh": emp1,
		"Hari":   emp2,
		"Ravi":   emp3,
	}

	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s--->Salary:%d  Country: %s\n", name, info.salary, info.country)
	}

}
