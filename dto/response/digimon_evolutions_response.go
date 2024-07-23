package response

type DigimonEvolutionsResponse struct {
	Number       int    `gorm:"column:number_base" json:"number"`
	Name         string `gorm:"column:name" json:"name"`
	Icon         string `gorm:"column:icon" json:"icon"`
	NextForm     int    `gorm:"column:next_form" json:"next_form"`
	NextFormName string `gorm:"column:next_form_name" json:"next_form_name"`
	NextFormIcon string `gorm:"column:next_form_icon" json:"next_form_icon"`
	Level        int    `gorm:"column:level" json:"level,omitempty"`
	ATK          int    `gorm:"column:atk" json:"atk,omitempty"`
	DEF          int    `gorm:"column:def" json:"def,omitempty"`
	HP           int    `gorm:"column:hp" json:"hp,omitempty"`
	INT          int    `gorm:"column:int" json:"int,omitempty"`
	SP           int    `gorm:"column:sp" json:"sp,omitempty"`
	SPD          int    `gorm:"column:spd" json:"spd,omitempty"`
	ABI          int    `gorm:"column:abi" json:"abi,omitempty"`
	CAM          int    `gorm:"column:cam" json:"cam,omitempty"`
	Additional   string `gorm:"column:additional" json:"additional,omitempty"`
}
