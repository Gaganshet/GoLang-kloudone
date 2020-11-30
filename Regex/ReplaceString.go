package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "C-Golang-Php-Python-Java-.Net"
	cpy := regexp.MustCompile("^(.*?)Php(.*)$")
	repStr := "${1}Java8$2"
	output := cpy.ReplaceAllString(str, repStr)
	fmt.Println(output)
}
