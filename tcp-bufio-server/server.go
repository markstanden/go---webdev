package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const (
	timeout = time.Second * 10
)

func main() {
	//create a tcp listener
	port := "8000"
	lisTCP, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Panicf("Error starting TCP server on port %v: %v", port, err)
	}
	// close listener when finished
	defer lisTCP.Close()

	//infinite loop, waiting for connections
	for {
		// wait for a connection
		tcp, err := lisTCP.Accept()
		if err != nil {
			log.Printf("Connection Error: %v", err)
		}

		// confirm connection
		io.WriteString(tcp, "\nConnection Established\n")
		
		// go routine to read incoming request
		go handle(tcp)
	}
}

func handle(tcpConn net.Conn) {
	err := tcpConn.SetDeadline(time.Now().Add(timeout))
	if err != nil {
		fmt.Println("Connection Timeout")
	}
	// bufio.NewScanner returns a pointer to a bufio.Scanner struct
	scanner := bufio.NewScanner(tcpConn)
	
	// scanner.Scan() advances the scanner to the next token (default is new line char)
	// returns true is successful, false if EOF
	for scanner.Scan() {
		line := scanner.Text() // could also use scanner.Bytes()
		fmt.Println(line) // print the line of text to the Stdout
		fmt.Fprintf(tcpConn, "Message Received")

	}

	// close the tcp connection
	defer func() {
		fmt.Fprintf(tcpConn, "Closing Connection")
		tcpConn.Close()
	}()
}