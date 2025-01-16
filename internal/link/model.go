package link

import (
	"math/rand"
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func Newlink(url string) *Link{
	link := &Link{
		Url: url,
	}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash(){
	link.Hash = Hash(10)
}

var letterRandom = []rune("abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ")

func Hash(n int) (string){
	b := make([]rune, n)
	
	for i := range b{
		b[i] = letterRandom[rand.Intn(len(letterRandom))]
	}

	return string(b)
}