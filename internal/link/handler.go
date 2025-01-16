package link

import (
	"fmt"
	"go/project_go/pkg/req"
	"go/project_go/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
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
		body, err := req.HandleBody[CreateLinkRequest](&w, r)
		if err != nil {
			return
		}
		link := Newlink(body.Url)

		for {
			existing_hash, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existing_hash == nil {
				break
			}

			link.GenerateHash()
		}

		resp, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.EncodeJson(w, resp, 201)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[UploadLinkRequest](&w, r)
		if err != nil {
			return
		}

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		link, err := handler.LinkRepository.Upload(&Link{
			Model: gorm.Model{
				ID: uint(id),
			},
			Url: body.Url,
			Hash: body.Hash,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.EncodeJson(w, link, 201)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, I'm Delete")
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		http.Redirect(w, r, link.Url, 200)
	}
}