package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("G([a-z]+)S")
	fmt.Println(re.FindStringSubmatch("GaganShet"))
	fmt.Println(re.FindStringSubmatch("Shet"))
}
