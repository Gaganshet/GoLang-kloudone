package main
import "fmt"

func main() {
	names := [3]string{"Mark Zuckerberg", "Bill Gates", "Larry Page"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}