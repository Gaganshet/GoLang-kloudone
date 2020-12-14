package main

import "fmt"

func main() {
	x := map[string]int{"A": 33, "B": 34, "C": 35, "D": 41, "E": 56, "F": 64, "G": 74, "H": 68, "I": 96, "J": 160}
	fmt.Println("The Map is :: ", x)
	fmt.Println("Slice from 1 to 9 :: ", x["H"])
}