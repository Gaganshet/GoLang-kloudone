package main

import (
	"fmt"
)

func main() {
	employeeSalary := make(map[string]int)
	employeeSalary["Ram"] = 50000
	employeeSalary["Ramu"] = 90000
	employeeSalary["Ramesh"] = 120000
	fmt.Println("employeeSalary map contents:", employeeSalary)
}
