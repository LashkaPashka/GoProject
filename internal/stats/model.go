package stats

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)


type Stat struct{
	*gorm.Model
	Id        uint `json:"id" gorm:"primarykey"`
	LinkId uint `json:"link_id"`
	Clicks int `json:"clicks"`
	Date datatypes.Date `json:"date"`
}