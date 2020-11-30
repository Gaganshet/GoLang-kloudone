package main

import "fmt"

type Student struct {
	name   string
	branch string
	batch  int
}

type Teacher struct {
	name    string
	subject string
	exp     int
	details Student
}

func main() {
	result := Teacher{
		name:    "Suresh",
		subject: "Java",
		exp:     1,
		details: Student{"Gagan", "CSE", 2020},
	}
	fmt.Println("Details of the Teacher")
	fmt.Println("Teacher's name: ", result.name)
	fmt.Println("Subject: ", result.subject)
	fmt.Println("Experience: ", result.exp)

	fmt.Println("\nDetails of Student")
	fmt.Println("Student's name: ", result.details.name)
	fmt.Println("Student's branch name: ", result.details.branch)
	fmt.Println("Year: ", result.details.batch)
}
