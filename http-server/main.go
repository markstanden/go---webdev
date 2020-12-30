package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

const port string = "8081"

func main() {
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalf("Server failed to start... \n%s", err.Error())
	} else {
		log.Printf("Listening on port :%s", port)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("Failed to make connection... \n%s", err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// do something for every line read in
		line := scanner.Text()
		fmt.Println(line)

		//parse http request info
		if i == 0 {
			req := strings.Fields(line)
			fmt.Println("----------------------")
			fmt.Printf("Method: %s\n", req[0])
			fmt.Printf("Route: %s\n", req[1])
			fmt.Printf("HTTP Version: %s\n", req[2])
			fmt.Println("----------------------")
		}

		if line == "" {
			// header is complete
			break
		}
		i++
	}
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

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprint(conn, body)
}
 