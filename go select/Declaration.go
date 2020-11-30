package main

import "fmt"

func main() {

	mychannel := make(chan int)
	select {
	case <-mychannel:

	default:
		fmt.Println("Not found")
	}
}
