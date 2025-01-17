package main

import (
	"fmt"
	"go/project_go/configs"
	"go/project_go/internal/auth"
	"go/project_go/internal/link"
	"go/project_go/pkg/db"
	"go/project_go/pkg/middleware"
	"net/http"
)

func main(){
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Respositories
	LinkRepository := link.NewLinkRepository(db)

	// Handler 
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	link.NetLinkHandler(router, link.LinkHandler{
		LinkRepository: LinkRepository,
	})
	
	server := http.Server{
		Addr: ":8081",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}