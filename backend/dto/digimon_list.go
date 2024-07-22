package dto

type DigimonList struct {
	Number    int    `gorm:"column:number" json:"number"`
	Name      string `gorm:"column:name" json:"name"`
	Stage     string `gorm:"column:stage" json:"stage"`
	Type      string `gorm:"column:type" json:"type"`
	Attribute string `gorm:"column:attribute" json:"attribute"`
	Image     string `gorm:"column:image" json:"image"`
	Icon      string `gorm:"column:icon" json:"icon"`
}
