package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("he")
	fmt.Println(re.FindStringIndex("hellomotto"))
	fmt.Println(re.FindStringIndex("heman"))
	fmt.Println(re.FindStringIndex("washe"))
}
