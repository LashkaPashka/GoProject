package main

import (
	"context"
	"fmt"
	"go/project_go/configs"
	"go/project_go/internal/auth"
	"go/project_go/internal/link"
	"go/project_go/internal/user"
	"go/project_go/pkg/db"
	"go/project_go/pkg/middleware"
	"net/http"
)

func main(){
	type key int

	const EmailKey key = 20
	
	ctx := context.WithValue(context.Background(), EmailKey, "a@a.ru")

	if value, ok := ctx.Value(EmailKey).(string); ok {
		fmt.Println(value)
	} else {
		fmt.Println("No value")
	}
}



func main1(){
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Respositories
	LinkRepository := link.NewLinkRepository(db)
	UserRepository := user.NewUserRepository(db)

	authService := auth.NewAuthService(UserRepository) 
	
	// Handler 
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
		AuthService: authService,
	})

	link.NetLinkHandler(router, link.LinkHandler{
		LinkRepository: LinkRepository,
	})
	

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr: ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}