package main
 
import (
    "fmt"
)
 
func main() {
    f := 13.8
    i := int(f)  // loses precision
    fmt.Println(i)     
     
    ii := 38
    ff := float64(ii)
     
    fmt.Println(ff)     
}