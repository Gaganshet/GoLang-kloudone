package main

import "fmt"

func main() {
	a := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)

	for index, value := range a {
		sum = sum + value
	}

	fmt.Printf("Sum of all the elements in array %v = %f", a, sum)
}
