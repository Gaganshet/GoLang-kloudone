package main 
  
import "fmt"
  
func main() { 
  
    // Creating an array 
    arr := [7]string{"This", "is", "the", "tutorial", 
                         "of", "Go", "language"} 
  
    // Display array 
    fmt.Println("Array:", arr) 
 
    myslice := arr[3: ] 
 
    fmt.Println("Slice:", myslice) 
    fmt.Printf("Length of the slice: %d", len(myslice)) 
    fmt.Printf("\nCapacity of the slice: %d", cap(myslice)) 
} 