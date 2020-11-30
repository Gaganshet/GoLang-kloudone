package main

import (
	"fmt"
)

func main() {
	employeeSalary := map[string]int{
		"Keshav": 12000,
		"Hari":   15000,
		"Kumar":  9000,
	}
	fmt.Println("map before deletion", employeeSalary)
	delete(employeeSalary, "Hari")
	fmt.Println("map after deletion", employeeSalary)

}
