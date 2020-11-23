package main

import "fmt"

type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {
	var a Address
	fmt.Println(a)

	a1 := Address{"Shanker", "Bangalore", 123478}
	fmt.Println("Address1: ", a1)

	a2 := Address{Name: "Teju", city: "sullia", Pincode: 7894331}
	fmt.Println("Address2: ", a2)

	a3 := Address{Name: "Chennai"}
	fmt.Println("Address3: ", a3)
}
