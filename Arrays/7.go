package main
import "fmt"

func main() {
	// Letting Go compiler infer the length of the array
	a := [...]int{3, 5, 7, 9, 11, 13, 17}
	fmt.Println(a)
}