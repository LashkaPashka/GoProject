package middleware

import "net/http"

type Wrapper struct {
	http.ResponseWriter
	statusCode int
}


func (w *Wrapper) WriteHeader(statusCode int){
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}