package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func root(res http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		f, fh, err := req.FormFile("q")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		
		fmt.Printf("\nFile:\t%v \nHeader:\t%v \nError:\t%v", f, fh, err)

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	res.Header().Set("Content-Type", "text/html; charset=UTF-8")
	io.WriteString(res, `
	<form method="POST" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="submit">
	</form>
	<br>
	` + s)

}