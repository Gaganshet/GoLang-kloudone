package main

import (
	"fmt"
	"strings"
	"bufio"
	"net"
	"os"
	
)

func main() {
	arg:= os.Args
	if len(arg) == 1 {
		fmt.Println("Enter Host Port Number")
		return
	}

	CONNECT := arg[1]
	fmt.Println("Client started on port number---->", CONNECT)
	fmt.Println("******READY TO CHAT*******\n")
	conn, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>Say Hello>>")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("--->: " + message)

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting")
			return
		}
	}
}
