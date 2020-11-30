package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "31/07/2020"
	str2 := "1/13/2020"
	str3 := "29/2/2020"
	str4 := "31/08/2020"
	str5 := "29/02/2020"
	str6 := "29/02/2020"
	str7 := "55/02/2030"
	str8 := "2_/02/2090"

	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Printf("\nDate: %v :%v\n", str1, re.MatchString(str1))
	fmt.Printf("Date: %v :%v\n", str2, re.MatchString(str2))
	fmt.Printf("Date: %v :%v\n", str3, re.MatchString(str3))
	fmt.Printf("Date: %v :%v\n", str4, re.MatchString(str4))
	fmt.Printf("Date: %v :%v\n", str5, re.MatchString(str5))
	fmt.Printf("Date: %v :%v\n", str6, re.MatchString(str6))
	fmt.Printf("Date: %v :%v\n", str7, re.MatchString(str7))
	fmt.Printf("Date: %v :%v\n", str8, re.MatchString(str8))
}
