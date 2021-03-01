package main 

import ( 
	"fmt"
) 

type Teacher struct{ 
	name string 
} 
func (t Teacher) Name() string {  
	return t.name
} 

type Student1 struct{ 
	Teacher 
} 

type Student2 struct{ 	
	Teacher 
} 

func main() { 
	 
	n1 := Student1{ 
	 
		Teacher{ 
				name: "Ganesh Shet", 
			}, 
		} 
		fmt.Println("Name is:", n1.Name())	 
	
	n2 := Student2{ 

		Teacher{ 
			name : "Shanker Shet", 
		}, 
	} 
	fmt.Println("Name is:", n2.Name()) 
} 
