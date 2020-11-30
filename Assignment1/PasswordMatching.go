package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "ç$/az@gmail.com"
	str2 := "ram@gmail_yahoo.com"
	str3 := "krishna@gmail-yahoo.com"
	str4 := "hari123@gmailyahoo"
	str5 := "om@gmail.yahoo"

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	fmt.Printf("Pattern: %v\n", re.String())
	fmt.Printf("\nEmail: %v :%v\n", str1, re.MatchString(str1))
	fmt.Printf("Email: %v :%v\n", str2, re.MatchString(str2))
	fmt.Printf("Email: %v :%v\n", str3, re.MatchString(str3))
	fmt.Printf("Email: %v :%v\n", str4, re.MatchString(str4))
	fmt.Printf("Email: %v :%v\n", str5, re.MatchString(str5))
}
