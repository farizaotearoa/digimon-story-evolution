package dto

type DigimonList struct {
	Number int    `gorm:"column:number" json:"number"`
	Name   string `gorm:"column:name" json:"name"`
	Stage  string `gorm:"column:stage" json:"stage"`
}
