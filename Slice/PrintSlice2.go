package main

import "fmt"

func main() {

	arr := [5]int{2, 3, 4, 8, 9}

	fmt.Println("Array:", arr)

	myslice := arr[:]

	fmt.Println("Slice:", myslice)
	fmt.Printf("Length of the slice: %d", len(myslice))
	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
}
