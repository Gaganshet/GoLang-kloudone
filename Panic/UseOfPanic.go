package main

import "fmt"

func entry(lang *string, aname *string) {
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
	entry(&A_lang, nil)
}
