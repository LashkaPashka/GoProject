package main

import (
	"fmt"
	"go/project_go/configs"
	"go/project_go/internal/auth"
	"go/project_go/pkg/db"
	"net/http"
)

func main(){
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	
	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}