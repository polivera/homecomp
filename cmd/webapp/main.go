package main

import (
	"context"
	"fmt"
	"net/http"

	"homecomp/pkg/configs"
	"homecomp/pkg/handlers/auth"
)

func main() {
	mux := http.NewServeMux()
	conf, err := configs.NewConfig()
	if err != nil {
		panic("cannot load configuration")
	}
	ctx := context.Background()

	auth.NewLoginHandler(conf, ctx).Handle(mux)

	mux.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	fmt.Println("Starting server on port 8008")
	http.ListenAndServe(":8008", mux)
}
