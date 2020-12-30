package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func main() {
	//create a tcp listener
	port := "8080"
	lisTCP, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Panic("Error starting TCP server on port %v: %v", port, err)
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

func handle(tcp net.Conn) {
	// bufio.NewScanner returns a pointer to a bufio.Scanner struct
	scanner := bufio.NewScanner(tcp)
	
	// scanner.Scan() advances the scanner to the next token (default is new line char)
	// returns true is successful, false if EOF
	for scanner.Scan() {
		line := scanner.Text() // could also use scanner.Bytes()
		log.Println(line) // print the line of text to the Stdout
	}

	// close the tcp connection
	defer tcp.Close();


}