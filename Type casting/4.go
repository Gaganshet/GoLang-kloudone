package main
 
import (
    "fmt"
)
 
func main() {
    var s string = "Hello World"
    var b []byte = []byte(s)     // convert ty bytes
     
    fmt.Println(b)  // [72 101 108 108 111 32 87 111 114 108 100]
     
    ss := string(b)              // convert to string
     
    fmt.Println(ss)     // Hello World
}