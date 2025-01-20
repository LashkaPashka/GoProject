package middleware

import (
	"context"
	"go/project_go/configs"
	"go/project_go/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	Emailkey key = "Emailkey"
)

func WriteStatusCode(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, conf *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer ") {
			WriteStatusCode(w)
			return
		}

		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJwt(conf.Auth.Secret).Parse(token)
		if !isValid {
			WriteStatusCode(w)
			return
		}

		ctx := context.WithValue(r.Context(), Emailkey, data.Email)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
