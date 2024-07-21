package models

type DigimonDetails struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Number      int    `gorm:"column:number" json:"number"`
	Name        string `gorm:"column:name" json:"name"`
	Stage       string `gorm:"column:stage" json:"stage"`
	Type        string `gorm:"column:type" json:"type"`
	Attribute   string `gorm:"column:attribute" json:"attribute"`
	Description string `gorm:"column:description" json:"description"`
	Image       string `gorm:"column:image" json:"image"`
	Icon        string `gorm:"column:icon" json:"icon"`
}

func (DigimonDetails) TableName() string {
	return "digimon_details"
}
