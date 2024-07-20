package models

type DigivolveRequirement struct {
	ID         int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Number     int    `gorm:"column:number" json:"number"`
	Level      int    `gorm:"column:level" json:"level"`
	ATK        int    `gorm:"column:atk" json:"atk"`
	DEF        int    `gorm:"column:def" json:"def"`
	HP         int    `gorm:"column:hp" json:"hp"`
	INT        int    `gorm:"column:int" json:"int"`
	SP         int    `gorm:"column:sp" json:"sp"`
	SPD        int    `gorm:"column:spd" json:"spd"`
	ABI        int    `gorm:"column:abi" json:"abi"`
	CAM        int    `gorm:"column:cam" json:"cam"`
	Additional string `gorm:"column:additional" json:"additional"`
}

func (DigivolveRequirement) TableName() string {
	return "digivolve_requirement"
}
