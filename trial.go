package main

import (
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseFiles("index.html"))

func index(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1>Kaleb Mesfin</h1><p>I am not sick of programming</p>"))
	templ.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	http.ListenAndServe(":8080", mux)

}
