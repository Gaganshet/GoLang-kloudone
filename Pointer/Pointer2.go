package main

import "fmt"

func main() {
   var a int = 20   /* actual variable declaration */
   var ip *int      /* pointer variable declaration */

   ip = &a  

   fmt.Printf("Address of a variable: %x\n", &a  )

   fmt.Printf("Address stored in ip variable: %x\n", ip )

   fmt.Printf("Value of *ip variable: %d\n", *ip )
}