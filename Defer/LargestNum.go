package main

import (
	"fmt"
)

func finished() {
	fmt.Println("Finished finding largest")
}
func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}
func main() {
	nums := []int{782, 1009, 2000, 563, 300}
	largest(nums)
}
