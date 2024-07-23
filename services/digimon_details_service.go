package services

import (
	"digimon-story-wiki/dto/request"
	"digimon-story-wiki/dto/response"
	"digimon-story-wiki/utils"
	"fmt"
	"go.uber.org/zap"
)

const (
	queryDigimonDetails = `SELECT 
			dd.number, 
			dd.name, 
			dd.stage, 
			dd.type, 
			dd.attribute, 
			dd.description,
			dd.image, 
			dd.icon
		FROM 
			digimon.digimon_details dd 
		WHERE
			dd.number = %d;`
	queryDigimonEvolutions = `SELECT
			dd.number as number_base,
			dd.name,
			dd.icon,
			de.next_form, 
			dt.name as next_form_name,
			dt.icon as next_form_icon,
			dr.*
		FROM
			digimon.digimon_details dd
		JOIN
			digimon.digimon_evolution de
		ON
			dd.number = de.number
		JOIN
			digimon.digimon_details dt
		ON
			de.next_form = dt.number
		JOIN
			digimon.digivolve_requirements dr 
		ON
			de.next_form = dr.number 
		WHERE
			de.number = %d or de.next_form = %d
		ORDER BY
			dd.number, de.next_form;`
)

func GetDigimonDetails(req request.DigimonDetailsRequest) (response.DigimonDetailsResponse, error) {
	var digimonDetails response.DigimonDetailsResponse
	query := fmt.Sprintf(queryDigimonDetails, req.Number)
	utils.Logger.Info(query)
	if err := utils.DB.Raw(query).Scan(&digimonDetails).Error; err != nil {
		utils.Logger.Error("Failed to GetDigimonDetails", zap.Int("number", req.Number), zap.Error(err))
		return response.DigimonDetailsResponse{}, err
	}
	return digimonDetails, nil
}

func GetDigimonEvolution(req request.DigimonDetailsRequest) ([]response.DigimonEvolutionsResponse, error) {
	var digimonEvolutions []response.DigimonEvolutionsResponse
	query := fmt.Sprintf(queryDigimonEvolutions, req.Number, req.Number)
	utils.Logger.Info(query)
	if err := utils.DB.Raw(query).Scan(&digimonEvolutions).Error; err != nil {
		utils.Logger.Error("Failed to GetDigimonEvolution", zap.Int("number", req.Number), zap.Error(err))
		return nil, err
	}
	return digimonEvolutions, nil
}
