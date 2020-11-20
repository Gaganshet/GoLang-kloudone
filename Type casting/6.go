package main

import (
    "fmt"
)

func main() {
    var a int = 28
    var f float64
    f=float64(a)
    fmt.Printf("%v,%T",f,f)
    fmt.Printf("%v,%T",a,a)  
}