package main
import (
	"fmt"
 )
  
 func main() {
  
	str1 := "abcdefg"
  
	// convert string to array of chars
	chars := []byte(str1)
  
	// print array of chars
	fmt.Println(chars)
  
	// convert char array to string
	str2 := string(chars)
  
	// print new string
	fmt.Println(str2)
  
 }