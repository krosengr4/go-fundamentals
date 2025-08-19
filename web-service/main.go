package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	// Handle the requests to the "/" path
	http.HandleFunc("/", Handler)
	// Using nil for the handler will use the DefaultServeMux provided by Go
	http.ListenAndServe(":3000", nil)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Using os.Open we can get the the menu.txt file
	f, _ := os.Open("./menu.txt")

	// io.Copy allows us to copy from a read source to a write source
	io.Copy(w, f)
}
