package main

import (
	"html/template"
	"net/http"
)

type Course struct {
	Title, Code, Description string
	ECTS                     int
}

var templ = template.Must(template.ParseFiles("index.html"))

func index(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1>Kaleb Mesfin</h1><p>I am not sick of programming</p>"))
	// templ.Execute(w, nil)
	course := Course{"DLD", "ITSE-3182", "Lorem Ipsum", 7}
	templ.Execute(w, course)
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", index)
	http.ListenAndServe(":8080", mux)

}
