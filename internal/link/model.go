package link

import (
	"go/project_go/internal/stats"
	"math/rand"

	"gorm.io/gorm"
)

type LinkAll struct{
	Link []Link `json:"links"`
	Count int64 `json:"count"`
}

type Link struct {
	gorm.Model
	Url string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
	Stats []stats.Stat `gorm:"constraints:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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