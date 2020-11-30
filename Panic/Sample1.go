package main

import (
	"fmt"
)

func entry(lang *string, aname *string) {
	defer fmt.Println("Defer statement in the entry function")
	if lang == nil {
		panic("Error: Language cannot be nil")
	}
	if aname == nil {
		panic("Error: Author name cannot be nil")
	}
	fmt.Printf("Author Language: %s \n Author Name: %s\n", *lang, *aname)
}

func main() {
	A_lang := "GO Language"
	defer fmt.Println("Defer statement in the Main function")
	entry(&A_lang, nil)
}
