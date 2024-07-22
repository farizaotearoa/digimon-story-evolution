package services

import (
	"digimon-story-wiki/dto"
	"digimon-story-wiki/utils"
	"go.uber.org/zap"
)

func GetAllDigimonList() ([]dto.DigimonList, error) {
	var digimon []dto.DigimonList
	query := `SELECT dd.number, dd.name, dd.stage, dd.type, dd.attribute, dd.image, dd.icon FROM digimon_details dd ORDER BY dd.number`
	if err := utils.DB.Raw(query).Scan(&digimon).Error; err != nil {
		utils.Logger.Error("Failed to GetAllDigimonList", zap.Error(err))
		return nil, err
	}
	return digimon, nil
}
