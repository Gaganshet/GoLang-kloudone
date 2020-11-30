package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "Hello X42 hai-3.35 string Z30"

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	fmt.Printf("Pattern: %v\n", re.String())
	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
