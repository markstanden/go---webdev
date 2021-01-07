package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type info struct{
	Title string;
	SubTitle string;
	Text string;
	ImagePath string;
}

var data []info

func main() {
	
	infoMain := info{"FileServer", "Using the FileServer method", "Here is a picture of something", "/index.jpeg"}
	data = append(data, infoMain)

	http.Handle("/", http.FileServer(http.Dir("../assets")))
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("../assets/index.gohtml")
	if err != nil {
		fmt.Println("Error loading html template", err)
	}
	tmp.Execute(res, data)
}

