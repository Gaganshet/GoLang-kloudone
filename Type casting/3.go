package main
 
import (
    "fmt"
)
 
func main() {
    f := 12.34567
    i := int(f)  // loses precision
    fmt.Println(i)      // 12
     
    ii := 34
    ff := float64(ii)
     
    fmt.Println(ff)     // 34
}