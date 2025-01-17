package auth

import (
	"errors"
	"go/project_go/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	user.UserRepository
}

func NewAuthService(repository *user.UserRepository) *AuthService{
	return &AuthService{
		UserRepository: *repository,
	}
}

func (service *AuthService) Login(email, password string) (string, error) {
	user, err := service.FindByEmail(email)
	if err != nil {
		return "", errors.New(ErrWrongCredetials)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredetials)
	}

	return user.Email, nil
}


func (service *AuthService) Register(email, password, name string) (string, error){
	exsiting_user, _ := service.UserRepository.FindByEmail(email)
	if exsiting_user != nil {
		return "", errors.New(UserExsiting)
	}
	hashedPasword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &user.User{
		Email: email,
		Name: name,
		Password: string(hashedPasword),
	}
	
	_, err = service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}


	return user.Email, nil
}