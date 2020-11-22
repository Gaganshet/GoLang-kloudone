package main
 
import (
    "fmt"
)
 
func main() {
    f := 13.8
    i := int(f)  // loses precision
    fmt.Println(i)     
     
    j := 38
    nn := float64(j)
     
    fmt.Println(nn)     
}