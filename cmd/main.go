package main

import (
	"net/http"
	"fmt"
	"go/project_go/configs"
	"go/project_go/internal/auth"
)

func main(){
	conf := configs.LoadConfig()
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