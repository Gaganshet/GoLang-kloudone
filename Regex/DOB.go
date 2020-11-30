package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "My date of birth is 1998-08-03"
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	fmt.Printf("Pattern: %v\n", re.String())
	fmt.Println(re.MatchString(str1))
	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
