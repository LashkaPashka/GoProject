package link

import (
	"math/rand"
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url string `json:"url"`
	Hash string `json:"hash" gorm:"unique"`
}

func Newlink(url string) *Link{
	return &Link{
		Url: url,
		Hash: GenerateHash(10),
	}
}

var letterRandom = []rune("abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateHash(n int) (string){
	b := make([]rune, n)

	for i := range b{
		b[i] = letterRandom[rand.Intn(len(letterRandom))]
	}

	return string(b)
}