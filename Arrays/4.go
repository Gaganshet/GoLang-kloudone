package main

import "fmt"

func main() {
	var x [5]int // An array of 5 integers

	x[0] = 100
	x[1] = 101
	x[3] = 103
	x[4] = 105

	fmt.Printf("The Arrays are:%d,%d,%d,%v\n ", x[0], x[1], x[3], x[4])
	fmt.Println("x = ", x)
}
