package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type info struct{
	Title string;
	SubTitle string;
	Text string;
	ImagePath string;
}

var data []info

func main() {
	
	infoMain := info{"FileServer", "Using the ServeContent method", "Here is a picture of something", "/index.jpeg"}
	data = append(data, infoMain)

	http.HandleFunc("/", index)
	http.HandleFunc("/index.jpeg", indexHero)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}

func index(resW http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("./index.gohtml")
	if err != nil {
		fmt.Println("Error loading html template", err)
	}
	tmp.Execute(resW, data)
}

func indexHero(resW http.ResponseWriter, req *http.Request) {
	file, err := os.Open("./index.jpeg")
	if err != nil {
		fmt.Println("Error opening File", err)
	}
	defer file.Close()

	f, err := file.Stat()
	if err != nil {
		fmt.Println("Error statting file", err)
	}

	fmt.Printf("File Info\nName:\t%v\nSize:\t%v\nModified:\t%v\n", f.Name(), f.Size(), f.ModTime())

	http.ServeContent(resW, req, f.Name(), f.ModTime(), file)
}

