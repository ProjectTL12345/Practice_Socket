package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	errorCheck(err)

	go func() {
		data := make([]byte, 4096)

		for {
			n, err := conn.Read(data)
			errorCheck(err)

			log.Println("Server send: " + string(data[:n]))
			time.Sleep(time.Duration(3) * time.Second)
		}
	}()

	for {
		var s string
		fmt.Scanln(&s)
		conn.Write([]byte(s))
		time.Sleep(time.Duration(3) * time.Second)
	}
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
