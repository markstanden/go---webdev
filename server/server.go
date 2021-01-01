package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var addr string = "localhost"
var port string = "8000"
var timeout = time.Second * 10

var wg sync.WaitGroup

func main() {
	var tcp net.Listener // TCP listener
	tcp, err := net.Listen("tcp", addr+":"+port)
	if err != nil {
		log.Fatalf("Error starting TCP listener : \n%v\n", err.Error())
	}

	// close the connection and release the port on close
	defer tcp.Close()

	for {
		var conn net.Conn
		conn, err = tcp.Accept()
		if err != nil {
			fmt.Println("Error opening TCP connection : \n", err.Error())
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Client Address : %v", conn.LocalAddr())

	// set a deadline and check for errors
	err := conn.SetDeadline(time.Now().Add(timeout))
	if err != nil {
		fmt.Println("Connection Timeout")
	}

	var scanner *bufio.Scanner

	scanner = bufio.NewScanner(conn)

	for scanner.Scan() {
		// From RFC 2616
		// first line of header METHOD<SPACE>ROUTE<SPACE>PROTOCOL VERSION<NEW-LINE>
		// HTTP Methods GET, HEAD, POST, PUT, DELETE, CONNECT, OPTIONS, TRACE, PATCH
		// HTTP-Version   = "HTTP" "/" 1*DIGIT "." 1*DIGIT
		currentLine := scanner.Text()

		
		fmt.Fprintln(conn, currentLine)



		if currentLine == "" {
			// end of header
			break
		}
	}

	respond(conn)
	log.Println("Response Sent...")
}

func respond(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title> Test Page </title>
		</head>
		<body>
			<h1>Test Page</h1>
			<p>Hello World!</p>
		</body>
	</html>`
	
	log.Println("Sending Response...")
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprint(conn, body)
	log.Println("Response Sent...")
}