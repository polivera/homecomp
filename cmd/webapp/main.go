package main

import (
	"fmt"
	"net/http"

	"homecomp/internal/configs"
	"homecomp/internal/database"
	"homecomp/pkg/handlers/auth"
)

func main() {
	mux := http.NewServeMux()
	conf, err := configs.NewConfig()
	if err != nil {
		panic("cannot load configuration")
	}
	db, err := database.NewConnection(conf.Database)
	if err != nil {
		panic("cannot connect to database")
	}

	auth.NewLoginHandler(conf, db).Handle(mux)

	mux.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	fmt.Println("Starting server on port 8008")
	http.ListenAndServe("localhost:8008", mux)
}
