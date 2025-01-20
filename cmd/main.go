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
	"time"
)

func Ticker(ctx context.Context){
	tick := time.NewTicker(1000*time.Millisecond)
	
	for {
		select{
			case <-tick.C:
				fmt.Println("Tick")
			case <-ctx.Done():
				fmt.Println("Cancel")	
				return
		}

	}

}


func main(){
	ctx, cancel := context.WithCancel(context.Background())
	go Ticker(ctx)

	time.Sleep(10*time.Second)
	cancel()
	time.Sleep(2*time.Second)

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