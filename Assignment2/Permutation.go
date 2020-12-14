package main

import (
	"fmt"
)

var result int

func main() {
	fmt.Print("Enter the value of n : ")
	var n int
	fmt.Scanln(&n)
	fmt.Print("Enter the value of r : ")
	var r int
	fmt.Scanln(&r)
	result := permutation(n, r)
	if result != 0 {
		fmt.Println(" ")
		fmt.Println("Permutation is:", result)
	}
}

func permutation(n int, r int) int {
	if n < r {
		fmt.Println(" ")
		fmt.Println("r should not be grater than n")
	} else {
		result = factorial(n) / factorial(n-r)
	}
	return result
}

func factorial(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f = f * i
	}
	return f
}
