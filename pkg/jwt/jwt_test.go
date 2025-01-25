package jwt_test

import (
	"go/project_go/pkg/jwt"
	"testing"
)

func TestCreateJWT(t *testing.T) {
	const email string = "test@test.ru"

	initial_jwt := jwt.NewJwt("xBRtjYStw9Gz3z4j6sJIiF//YU12swBa0glAw7h3VUk=")
	
	token, err := initial_jwt.CreateJWT(email)
	if err != nil {
		t.Fatal(err)
	}
	
	isValid, data := initial_jwt.Parse(token)
	if !isValid {
		t.Fatalf("Token not valid")
	}

	if email != data.Email {
		t.Fatalf("Email %s not match %s", data.Email, email)
	}
}