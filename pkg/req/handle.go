package req

import (
	"go/project_go/pkg/res"
	"net/http"
)


func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error){
	body, err := Decode[T](r.Body)
	if err != nil {
		res.EncodeJson(*w, err.Error(), 409)
		return nil, err
	}

	err = isValid(body)
	if err != nil {
		res.EncodeJson(*w, err.Error(), 422)
		return nil, err
	}

	return &body, nil
}