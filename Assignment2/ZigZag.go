package main

import "fmt"

func main() {
	arr := [][]int{{10, 20, 30, 40}, {15, 86, 77, 87}, {91, 80, 8, 44}, {13, 14, 14, 17}, {55, 44, 99, 44}}
	fmt.Println("The Array is :", arr)
	fmt.Println(" ")
	fmt.Println("ZigZag Matrix is : ")
	zigzagMatrix(arr)
}

func min2(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func min3(a int, b int, c int) int {
	return min2(min2(a, b), c)
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func zigzagMatrix(matrix [][]int) {
	for line := 1; line <= 8; line++ {
		start_col := max(0, line-5)
		count := min3(line, (4 - start_col), 5)
		for j := 0; j < count; j++ {
			fmt.Print(matrix[min2(5, line)-j-1][start_col+j], " ")
		}
		fmt.Println(" ")
	}
}
