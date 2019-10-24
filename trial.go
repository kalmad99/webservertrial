package main

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	http.ListenAndServe(":8080", mux)

}
