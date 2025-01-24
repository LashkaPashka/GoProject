package main

import (
	"go/project_go/internal/link"
	"go/project_go/internal/stats"
	"go/project_go/internal/user"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	db.AutoMigrate(&link.Link{}, &user.User{}, &stats.Stat{})
}