package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting...")

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		panic(err)
	}

	fmt.Println("Past the listen... :)")

	for {
		fmt.Println("about to accept connections")
		conn, err := l.Accept()
		fmt.Println("i've accepted connections")
		if err != nil {
			panic(err)
		}

		fmt.Println("ain't gonna do anything")

		go func(c net.Conn) {
			fmt.Println("writing nothing")
			c.Write([]byte{})
		}(conn)
	}
}