package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

var addr string = "localhost"
var port string = "8080"

func main() {
	// creates a Dialer and runs the dial method on it, returning a connection and an error
	tcpConn, err := net.Dial("tcp", addr + ":" + port)
	if err != nil {
		log.Panicf("Failed to connect to %v:%v \n%v\n", addr, port, err)
	}
	// close connection once complete
	defer tcpConn.Close()

	bs, err := ioutil.ReadAll(tcpConn)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs))


}
