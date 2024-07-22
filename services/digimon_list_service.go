package services

import (
	"digimon-story-wiki/dto/request"
	"digimon-story-wiki/dto/response"
	"digimon-story-wiki/utils"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

const (
	queryDigimonList = `SELECT 
			dd.number, 
			dd.name, 
			dd.stage, 
			dd.type, 
			dd.attribute, 
			dd.image, 
			dd.icon 
		FROM 
			digimon.digimon_details dd 
		%s 
		ORDER BY 
			%s %s 
		LIMIT %d OFFSET %d;`
	queryDigimonListSize = `SELECT 
			COUNT(*) 
		FROM 
			digimon.digimon_details dd 
		%s;`
)

var sortBy = map[string]string{
	"number": "dd.number",
	"name":   "dd.name",
	"stage":  "array_position(ARRAY['Training 1', 'Training 2', 'Rookie', 'Champion', 'Ultimate', 'Mega', 'Ultra', 'Armor', 'No Stage'], dd.stage), dd.number",
}

func GetAllDigimonList(req request.DigimonListRequest) ([]response.DigimonListResponse, error) {
	var digimon []response.DigimonListResponse
	offset := (req.PageNum - 1) * req.PageSize
	query := fmt.Sprintf(queryDigimonList, constructWhereCondition(req), sortBy[req.SortBy], strings.ToUpper(req.SortOrder), req.PageSize, offset)
	utils.Logger.Info(query)
	if err := utils.DB.Raw(query).Scan(&digimon).Error; err != nil {
		utils.Logger.Error("Failed to GetAllDigimonList", zap.Error(err))
		return nil, err
	}
	return digimon, nil
}

func GetAllDigimonListSize(req request.DigimonListRequest) (int, error) {
	var count int
	query := fmt.Sprintf(queryDigimonListSize, constructWhereCondition(req))
	utils.Logger.Info(query)
	if err := utils.DB.Raw(query).Scan(&count).Error; err != nil {
		utils.Logger.Error("Failed to GetAllDigimonListSize", zap.Error(err))
		return 0, err
	}
	return count, nil
}

func constructWhereCondition(req request.DigimonListRequest) string {
	whereCondition := ""

	var inCondition []string
	inCondition = append(inCondition, constructInCondition(req.Stage, "dd.stage"))
	inCondition = append(inCondition, constructInCondition(req.Type, "dd.type"))
	inCondition = append(inCondition, constructInCondition(req.Attribute, "dd.attribute"))

	if inCondition[0] != "" || inCondition[1] != "" || inCondition[2] != "" {
		var nonEmpty []string
		for _, el := range inCondition {
			if el != "" {
				nonEmpty = append(nonEmpty, el)
			}
		}
		whereCondition = "WHERE " + strings.Join(nonEmpty, " AND ")
	}

	return whereCondition
}

func constructInCondition(value []string, column string) string {
	inCondition := ""
	for i, level := range value {
		value[i] = fmt.Sprintf("'%s'", level)
	}
	if len(value) > 0 {
		inCondition = fmt.Sprintf("%s IN (%s)", column, strings.Join(value, ","))
	}

	return inCondition
}
