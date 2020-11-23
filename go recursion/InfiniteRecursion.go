package main

import (
	"fmt"
)

func print_hello() {
	fmt.Println("hellohai")
	print_hello()
}
func main() {
	print_hello()
}
