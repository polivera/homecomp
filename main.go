package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", peperino)

	fmt.Println("Starting server on port 8008")
	http.ListenAndServe(":8008", mux)
}

func peperino(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "This is some %s data", "shitty in function")
}
