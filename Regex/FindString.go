package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("mm$")
	fmt.Println(re.FindString("Ramm"))
}
