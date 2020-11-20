package main

import "fmt"

func main() {
   var a int
   var ptr *int
   a = 3000

   /* take the address of var */
   ptr = &a

   /* take the value using pptr */
   fmt.Printf("Value of a = %d\n", a )
   fmt.Printf("Value available at *ptr = %d\n", *ptr )
}