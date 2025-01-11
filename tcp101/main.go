package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	//storing / readin the req
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	fmt.Printf(string(buff))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello,World\r\n"))
	fmt.Println("req severd")
	conn.Close()
}

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		// it will cill th procees
		log.Fatal(err)
	}
	for {
		//contyinoulsy accestion the client
		conn, err := listner.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go do(conn)
	}
}
