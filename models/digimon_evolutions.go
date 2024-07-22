package models

type DigimonEvolutions struct {
	ID       int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Number   int `gorm:"column:number" json:"number"`
	NextForm int `gorm:"column:next_form" json:"next_form"`
}

func (DigimonEvolutions) TableName() string {
	return "digimon_evolutions"
}
