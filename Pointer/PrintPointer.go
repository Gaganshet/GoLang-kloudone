package main

import "fmt"

func main() {
	var a int
	var ptr *int
	a = 3000

	ptr = &a

	fmt.Printf("Value of a = %d\n", a)
	fmt.Printf("Value available at *ptr = %d\n", *ptr)
}
