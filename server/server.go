package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

var addr string = "localhost"
var port string = "8000"
var timeout = time.Second * 10

func main() {
	var tcp net.Listener // TCP listener
	tcp, err := net.Listen("tcp", addr + ":" + port)
	if err != nil {
		log.Fatalf("Error starting TCP listener : \n", err.Error())
	}

	var conn net.Conn
	conn, err := tcp.Accept()
	if err != nil {
		fmt.Println("Error opening TCP connection : \n", err.Error())
	}

	go handleConnection(conn)


	// close the connection and release the port.
	defer tcp.Close()
	defer conn.Close()
}

func handleConnection(conn net.Conn) {
	fmt.Printf("Client Address : %v", conn.LocalAddr())

	var bs []byte

	conn.SetDeadline(time.Now() + timeout)
	_, err := conn.Read(bs)
	if err != nil {
		fmt.Println("Error reading connection string : \n", err.Error())
	}

}