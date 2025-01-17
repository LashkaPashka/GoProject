package auth

import (
	"fmt"
	"go/project_go/configs"
	"go/project_go/pkg/res"
	"net/http"
	"go/project_go/pkg/req"
)

type AuthHandlerDeps struct{
	*configs.Config
	*AuthService
}

type AuthHandler struct{
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps){
	handler := &AuthHandler{
		Config: deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		fmt.Println(handler.Config.Auth.Secret)
		
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		data := LoginResponse{
			Token: "333",
		}
		fmt.Println(body)

		email, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			return
		}
		
		fmt.Println(email)
		res.EncodeJson(w, data, 201)

	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}

		handler.AuthService.Register(body.Email, body.Password, body.Name)
	}
}