package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"text/template"
	"time"
)

var addr string = "localhost"
var port string = "8081"
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

	firstLine := true
	var htmlFileName string

	for scanner.Scan() {
		// From RFC 2616
		// first line of header METHOD<SPACE>ROUTE<SPACE>PROTOCOL VERSION<NEW-LINE>
		// HTTP Methods GET, HEAD, POST, PUT, DELETE, CONNECT, OPTIONS, TRACE, PATCH
		// HTTP-Version   = "HTTP" "/" 1*DIGIT "." 1*DIGIT
		currentLine := scanner.Text()
		if firstLine {
			requestLine := strings.Fields(currentLine)

			fmt.Fprintf(os.Stdout, "\r\n*************************\r\n")
			fmt.Fprintf(os.Stdout, "Method: %v\r\n", requestLine[0])
			fmt.Fprintf(os.Stdout, "Route: %v\r\n", requestLine[1])
			fmt.Fprintf(os.Stdout, "Protocol: %v\r\n", requestLine[2])
			fmt.Fprintf(os.Stdout, "*************************\r\n")

			if requestLine[0] == "GET" {

			switch requestLine[1] {
			case "/":
				htmlFileName = "./index.gohtml"
			case "/contact": 
				htmlFileName = "./contact.gohtml"
			}
		}
		
		} else {
			fmt.Fprintln(os.Stdout, currentLine)
		}

		if currentLine == "" {
			// end of header
			break
		}

		firstLine = false
	}

	if htmlFileName != "" {
		tplt, err := template.ParseFiles(htmlFileName)
		if err != nil {
			log.Panicln("Cannot load html :", err)
		}
		var bb bytes.Buffer
		tplt.Execute(&bb, nil)
		respond(conn, bb.String())
	}
	htmlFileName = ""
}

func respond(conn net.Conn, body string) {
	log.Println("Sending Response...")

	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprint(conn, body)
	log.Println("Response Sent...")
}
