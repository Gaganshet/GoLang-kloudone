package main 
  
import "fmt"
  
func main() { 
 
    arr := [9]int16{2,3,4,8,9,4,7,8,4} 
  
    fmt.Println("Array:", arr) 
 
    myslice := arr[7:] 
    fmt.Println( myslice) 
 }