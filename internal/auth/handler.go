package auth

import (
	"fmt"
	"go/project_go/configs"
	"go/project_go/pkg/jwt"
	"go/project_go/pkg/req"
	"go/project_go/pkg/res"
	"net/http"
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
		
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}

		email, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJwt(handler.Config.Auth.Secret).CreateJWT(email)
		fmt.Println(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		data := LoginResponse{
			Token: token,
		}

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
		
		token, err := jwt.NewJwt(handler.Auth.Secret).CreateJWT(body.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return 
		}

		data := LoginResponse{
			Token: token,
		}
		res.EncodeJson(w, data, 201)
	}
}