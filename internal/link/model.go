package link

import (
	"gorm.io/gorm"
)


type Link struct {
	gorm.Model
	Url string `json:"url"`
	Hash string `json:"hash"`
}
