package link

import (
	"fmt"
	"go/project_go/pkg/req"
	"go/project_go/pkg/res"
	"net/http"
)


type LinkHandler struct{
	LinkRepository *LinkRepository
}



func NetLinkHandler(router *http.ServeMux, deps LinkHandler){
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}

	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())
}

func (handler *LinkHandler) Create() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkRequest](&w, r)
		if err != nil {
			return
		}
		url := Newlink(body.Url)
		
		resp, err := handler.LinkRepository.Create(url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.EncodeJson(w, resp, 200)
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