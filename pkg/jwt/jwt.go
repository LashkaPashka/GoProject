package jwt

import (
    "github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}

func NewJwt(secret string) *JWT{
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) CreateJWT(email string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": email,
		})
	sToken, err := token.SignedString(j.Secret)
	if err != nil {
		return "", err
	}

	return sToken, nil
}