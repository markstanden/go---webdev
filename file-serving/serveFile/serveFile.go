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
	
	infoMain := info{"FileServer", "Using the ServeContent method", "Here is a picture of something", "/assets/index.jpeg"}
	data = append(data, infoMain)

	http.HandleFunc("/", index)
	http.HandleFunc("/assets/index.jpeg", indexHero)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}

func index(resW http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("../assets/index.gohtml")
	if err != nil {
		fmt.Println("Error loading html template", err)
	}
	tmp.Execute(resW, data)
}

func indexHero(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "../assets/index.jpeg")
}

