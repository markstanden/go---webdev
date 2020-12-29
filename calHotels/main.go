package main

import (
	"os"
	"text/template"
)

var tmpl *template.Template

var hotelList []hotel 

type hotel struct{
	Name string
	Address string
	City string
	Zip int
	Region int
}

// Regions within california for hotels to be based
const (
	Southern = iota
	Central = iota
	Northern = iota
)

func (h hotel) GetRegion(reg int) string {
	switch {
	case reg == Southern:
		return "Southern"
	case reg == Central:
		return "Central"
	case reg == Northern:
		return "Northern"
	} 
	return ""
} 

func init() {
	tmpl = template.Must(template.ParseFiles("./index.gohtml"))
}

func main() {
hotelList = []hotel{
	hotel{Name: "Big", Address: "1 Big Road", City: "Reno", Zip: 13434, Region: Central},
	hotel{Name: "Bigger", Address: "1 10 lane wide Road", City: "San Fran", Zip: 22434, Region: Central},
	hotel{Name: "Small", Address: "1 Congested Road", City: "Fresno", Zip: 124134, Region: Southern},
	hotel{Name: "Greatest Hotel", Address: "1 Busy Road", City: "Oakland", Zip: 124344, Region: Southern},
	hotel{Name: "Worst Hotel", Address: "1 Massive Road", City: "San Diego", Zip: 1223434, Region: Northern},
	hotel{Name: "Tiny Hotel", Address: "1 Big Ass Road", City: "Silicon Valley", Zip: 142434, Region: Central},
}


	tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", hotelList)
}