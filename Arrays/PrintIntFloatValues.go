package main

import "fmt"

func main() {
	a := [2][2]int{
		{3, 5},
		{7, 9},
	}
	fmt.Println(a)

	b := [3][4]float64{
		{1, 3},
		{4.5, -3, 7.4, 2},
		{6, 2, 11},
	}
	fmt.Println(b)
}
