package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	//create a tcp listener
	port := "8080"
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Printf("Error starting TCP server on port %v: %v", port, err)
	}
	// close listener when finished
	defer lis.Close()

	//infinite loop, waiting for connections
	for {
		// wait for a connection
		con, err := lis.Accept()
		if err != nil {
			log.Printf("Connection Error: %v", err)
		}

		// output over connection
		io.WriteString(con, "\nConnection Established\n")
		fmt.Fprintln(con, "Test")
		
	}

}