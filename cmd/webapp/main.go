package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"homecomp/internal/configs"
	"homecomp/internal/database"
	"homecomp/internal/repositories"
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
		panic(fmt.Sprintf("cannot connect to database: %s", err.Error()))
	}

	userRepo := repositories.NewUserRepo(db)

	handlers.NewLoginHandler(conf, userRepo).Handle(mux)

	mux.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	serverAddress := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)

	srv := &http.Server{
		Addr:    serverAddress,
		Handler: mux,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	fmt.Printf("Starting server on %s\n", serverAddress)
	err = srv.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
