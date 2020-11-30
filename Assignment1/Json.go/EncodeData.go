package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Id         int
	Name       string
	Occupation string
}

func main() {
	u1 := User{1, "Ram", "Doctor"}
	json_data, err := json.Marshal(u1)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json_data))
	users := []User{
		{Id: 2, Name: "Roy", Occupation: "Teacher"},
		{Id: 3, Name: "Krishna", Occupation: "Dancer"},
		{Id: 4, Name: "Hari", Occupation: "Scientist"},
	}

	json_data2, err := json.Marshal(users)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json_data2))
}
