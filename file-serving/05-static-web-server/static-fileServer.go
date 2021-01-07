package main

import (
	"fmt"
	"net/http"
)
func main() {
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
