package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HellowServer)
	http.ListenAndServe(":80", nil)
}

func HellowServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>Hello</h1></body></html>")
}
