package main

import "fmt"

func main() {
	names := [3]string{"krishna", "ram", "balram"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
