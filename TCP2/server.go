package main

import (
	"fmt"
	"strings"
	"bufio"
	"net"
	"os"
	
)

func main() {
	var queue []string

	arg:= os.Args
	if len(arg) == 1 {
		fmt.Println("Enter port number")
		return
	}

	PORT := ":" + arg[1]
	fmt.Println("Server Started On Port Number===>",PORT)
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		queue = enqueue(queue, netData)
		fmt.Println(queue)

		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("---> ", string(netData))
		conn.Write([]byte(netData))
	}
}

func enqueue(queue []string, element string) []string {
	queue = append(queue, element)
	fmt.Println("Enqueue:", element)
	return queue
}
