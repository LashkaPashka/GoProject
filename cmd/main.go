package main

import (
	"fmt"
	"go/project_go/configs"
	"go/project_go/internal/auth"
	"go/project_go/internal/link"
	"go/project_go/internal/stats"
	"go/project_go/internal/user"
	"go/project_go/pkg/db"
	"go/project_go/pkg/event"
	"go/project_go/pkg/middleware"
	"net/http"
)

func App() http.Handler {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// EventBus
	EventBus := event.NewEventBus()

	// Respositories
	LinkRepository := link.NewLinkRepository(db)
	UserRepository := user.NewUserRepository(db)
	StatRepository := stats.NewStatsRepository(db)

	// Services
	authService := auth.NewAuthService(UserRepository) 
	serviceStat := stats.NewServiceStat(&stats.ServiceStatDeps{
		EventBus: EventBus,
		StatRepository: StatRepository,
	})

	// Handler 
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
		AuthService: authService,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: LinkRepository,
		Config: conf,
		EventBus: EventBus,
	})
	
	stats.NewStatHandler(router, stats.StatHandlerDeps{
		StatRepository: StatRepository,
		Config: conf,
	})

	go serviceStat.AddClick()

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	return stack(router)
}

func main(){
	app := App()

	server := http.Server{
		Addr: ":8081",
		Handler: app,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}