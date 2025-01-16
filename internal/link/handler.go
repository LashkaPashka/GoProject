package link

import (
	"fmt"
	"go/project_go/configs"
	"net/http"

)

type LinkHandlerDeps struct{
	*configs.Config
}

type LinkHandler struct{
	*configs.Config
}



func NetLinkHandler(router *http.ServeMux, deps LinkHandlerDeps){
	handler := &LinkHandler{
		Config: deps.Config,
	}

	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("POST /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{alias}", handler.GoTo())
}

func (handler *LinkHandler) Create() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, I'm Create")
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, I'm Update")
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, I'm Delete")
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, I'm GoTo")
	}
}