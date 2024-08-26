package main

import (
	"fmt"
	"net/http"

	"homecomp/internal/configs"
	"homecomp/internal/database"
	"homecomp/pkg/web/handlers"
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

	handlers.NewLoginHandler(conf, db).Handle(mux)

	mux.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	serverAddress := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	fmt.Printf("Starting server on %s\n", serverAddress)
	err = http.ListenAndServe(serverAddress, mux)
	if err != nil {
		panic(err.Error())
	}
}
