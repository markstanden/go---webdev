package main

import (
	"html/template"
	"os"
)

var tmpl *template.Template

type personalInfo struct {
	Name string
	Age  string
	Job  string
}

func init() {
	tmpl = template.Must(template.ParseGlob("./html/*.gohtml"))
}

func main() {
	mark := personalInfo{
		Name: "Mark",
		Age:  "40",
		Job:  "Electrician",
	}
	kelly := personalInfo{"Kelly", "42", "CEO Badger & Bear"}
	charlie := personalInfo{"Charlie", "8", "Developer"}
	joanie := personalInfo{"Joanie", "5", "Singer/Songwriter"}
	benji := personalInfo{"Benji", "2", "Demolitions Expert"}

	data := []personalInfo{mark, kelly, charlie, joanie, benji}

	tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
}
