package main

import (
	"context"
	"fmt"
	"net/http"

	"homecomp/internal/configs"
	"homecomp/internal/database"
	"homecomp/pkg/web/handlers"
)

func main() {
	conf, err := configs.NewConfig()
	ctx, cancel := context.WithTimeout(context.Background(), conf.App.Timeout)
	defer cancel()

	mux := http.NewServeMux()
	if err != nil {
		panic(fmt.Sprintf("cannot load configuration: %s", err.Error()))
	}
	db, err := database.NewConnection(conf.Database)
	if err != nil {
		panic("cannot connect to database")
	}

	handlers.NewLoginHandler(ctx, conf, db).Handle(mux)

	mux.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	serverAddress := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)

	fmt.Printf("Starting server on %s\n", serverAddress)
	err = http.ListenAndServe(serverAddress, mux)
	if err != nil {
		panic(err.Error())
	}
}
