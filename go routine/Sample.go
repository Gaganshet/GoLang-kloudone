package main

import "fmt"

func display(str string) {
	for w := 0; w < 4; w++ {
		fmt.Println(str)
	}
}

func main() {
	go display("Welcome")
	display("Happy Deepavali")
}
