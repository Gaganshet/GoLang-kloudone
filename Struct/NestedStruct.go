package main

import "fmt"

type Author struct {
	name   string
	branch string
	year   int
}

type HR struct {
	details Author
}

func main() {
	result := HR{
		details: Author{"Arjun", "CS", 2020},
	}
	fmt.Println("\nDetails of Author")
	fmt.Println(result)
}
