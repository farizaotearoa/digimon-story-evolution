package services

import (
	"context"
	"digimon-story-evolution/dto/request"
	"digimon-story-evolution/dto/response"
	"digimon-story-evolution/utils"

	"go.uber.org/zap"
)

// GetAllDigimonList fetches digimon list from Digimon Service
func GetAllDigimonList(ctx context.Context, req request.DigimonListRequest) ([]response.DigimonListResponse, error) {
	var digimonList []response.DigimonListResponse

	// Make request to Digimon Service
	resp, err := Clients.DigimonClient.client.Post(ctx, "/api/v1/digimon/list", req, nil)
	if err != nil {
		utils.Logger.Error("Failed to call Digimon Service", zap.Error(err))
		return nil, err
	}

	// Handle response
	if err := handleServiceResponse(resp, &digimonList); err != nil {
		return nil, err
	}

	return digimonList, nil
}

// GetAllDigimonListSize fetches digimon list size from Digimon Service
func GetAllDigimonListSize(ctx context.Context, req request.DigimonListRequest) (int, error) {
	var sizeResponse struct {
		Size int `json:"size"`
	}

	// Make request to Digimon Service
	resp, err := Clients.DigimonClient.client.Post(ctx, "/api/v1/digimon/list/size", req, nil)
	if err != nil {
		utils.Logger.Error("Failed to call Digimon Service for size", zap.Error(err))
		return 0, err
	}

	// Handle response
	if err := handleServiceResponse(resp, &sizeResponse); err != nil {
		return 0, err
	}

	return sizeResponse.Size, nil
}

// GetDigimonDetails fetches digimon details from Digimon Service
func GetDigimonDetails(ctx context.Context, req request.DigimonDetailsRequest) (*response.DigimonDetailsResponse, error) {
	var details response.DigimonDetailsResponse

	// Make request to Digimon Service
	resp, err := Clients.DigimonClient.client.Post(ctx, "/api/v1/digimon/details", req, nil)
	if err != nil {
		utils.Logger.Error("Failed to call Digimon Service for details", zap.Error(err))
		return nil, err
	}

	// Handle response
	if err := handleServiceResponse(resp, &details); err != nil {
		return nil, err
	}

	return &details, nil
}

// GetDigimonEvolution fetches digimon evolution from Digimon Service
func GetDigimonEvolution(ctx context.Context, req request.DigimonDetailsRequest) (*response.DigimonEvolutionsResponse, error) {
	var evolution response.DigimonEvolutionsResponse

	// Make request to Digimon Service
	resp, err := Clients.DigimonClient.client.Post(ctx, "/api/v1/digimon/evolutions", req, nil)
	if err != nil {
		utils.Logger.Error("Failed to call Digimon Service for evolution", zap.Error(err))
		return nil, err
	}

	// Handle response
	if err := handleServiceResponse(resp, &evolution); err != nil {
		return nil, err
	}

	return &evolution, nil
}
