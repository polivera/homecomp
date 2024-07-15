package main

import (
	"context"
	"fmt"
	"net/http"

	"gitlab.com/xapitan/homecomp/pkg/templates"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", peperino)

	fmt.Println("Starting server on port 8008")
	http.ListenAndServe(":8008", mux)
}

func peperino(w http.ResponseWriter, _ *http.Request) {
	innerContent := templates.Testing("Broter")
	component := templates.Layout(innerContent, "This is some good title", "en")
	component.Render(context.Background(), w)
}
