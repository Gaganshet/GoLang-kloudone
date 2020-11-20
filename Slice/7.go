package main 
  
import "fmt"
  
func main() { 
 
    arr := [5]int16{2,3,4,8,9} 
  
    fmt.Println("Array:", arr) 
 
    myslice := arr[0:4] 
 
    fmt.Println("Slice:", myslice) 
    fmt.Printf("Length of the slice: %d", len(myslice)) 
    fmt.Printf("\nCapacity of the slice: %d", cap(myslice)) 
} 