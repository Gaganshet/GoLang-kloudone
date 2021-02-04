
package main 
  
import ( 
    "fmt"
) 
  
type G1 struct{ 
 
    name string 
} 
func (g G1) print() string{ 
       
    return g.name
} 

type G2 struct{ 
  
    name string 
} 
 
func (g G2) print2() string{ 

    return g.name
} 
  
type child struct{ 
    G1
    G2 
} 

func main() { 
      
    c1 := child{ 
    
        G1{     
            name: "Hello", 
        }, 
        G2{ 
            name: "\nHow are you.\n", 
        }, 
    } 
    fmt.Println(c1.print()) 
    fmt.Println(c1.print2()) 
} 