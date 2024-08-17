package main

import (
	"context"
	"fmt"
	"homecomp/pkg/handlers/auth"
	"homecomp/pkg/templates"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	loginHandler := auth.NewLoginHandler()

	loginHandler.Handle(mux)
	mux.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	mux.HandleFunc("GET /hello", peperino)

	fmt.Println("Starting server on port 8008")
	http.ListenAndServe("localhost:8008", mux)
}

func peperino(w http.ResponseWriter, _ *http.Request) {
	innerContent := templates.Testing("Broter")
	component := templates.Layout(innerContent)
	component.Render(context.Background(), w)
}
