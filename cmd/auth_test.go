package main

import (
	"bytes"
	"encoding/json"
	"go/project_go/internal/auth"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"go/project_go/internal/user"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB{
	err := godotenv.Load(".env")
	if err != nil{
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func CreateTestDb(db *gorm.DB) {
	db.Create(&user.User{
		Email: "test@test.com",
		Password: "$2a$10$T3xP7bcpncGnAYBmv7z/beHf9NewkjIBBSMMiHE.Xjl2BJPBjeESi",
		Name: "Jessica",
	})
}

func RemoveTestDb(db *gorm.DB){
	db.Unscoped().Delete(&user.User{}, "email = ?", "test@test.com")
}

func TestLogin(t *testing.T) {
	db := InitDb()
	CreateTestDb(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email: "test@test.com",
		Password: "12345",
	})

	res, err := http.Post(ts.URL + "/auth/login", "application/json", bytes.NewReader(data))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 201 {
		t.Fatalf("Not %d and %d", 201, res.StatusCode)
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse

	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatalf("Not token")
	}

	RemoveTestDb(db)
}

func TestLoginFail(t *testing.T) {
	db := InitDb()
	CreateTestDb(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email: "test@test.com",
		Password: "1",
	})

	res, err := http.Post(ts.URL + "/auth/login", "application/json", bytes.NewReader(data))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 401 {
		t.Fatalf("Not %d and %d", 401, res.StatusCode)
	}
	RemoveTestDb(db)
}