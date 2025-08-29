package services

import (
	"context"
	"digimon-story-evolution/utils"

	"go.uber.org/zap"
)

// UserContribution represents a user's contribution
type UserContribution struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	DigimonID   string `json:"digimon_id"`
	Type        string `json:"type"` // "new", "update", "correction"
	Description string `json:"description"`
	Status      string `json:"status"` // "pending", "approved", "rejected"
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// CreateContribution creates a new user contribution
func CreateContribution(ctx context.Context, contribution UserContribution) (*UserContribution, error) {
	var createdContribution UserContribution

	// Make request to User Service
	resp, err := Clients.UserClient.client.Post(ctx, "/api/v1/contributions", contribution, nil)
	if err != nil {
		utils.Logger.Error("Failed to call User Service for creating contribution", zap.Error(err))
		return nil, err
	}

	// Handle response
	if err := handleServiceResponse(resp, &createdContribution); err != nil {
		return nil, err
	}

	return &createdContribution, nil
}

// GetUserContributions fetches user contributions
func GetUserContributions(ctx context.Context, userID string) ([]UserContribution, error) {
	var contributions []UserContribution

	// Make request to User Service
	resp, err := Clients.UserClient.client.Get(ctx, "/api/v1/users/"+userID+"/contributions", nil)
	if err != nil {
		utils.Logger.Error("Failed to call User Service for getting contributions", zap.Error(err))
		return nil, err
	}

	// Handle response
	if err := handleServiceResponse(resp, &contributions); err != nil {
		return nil, err
	}

	return contributions, nil
}

// GetContribution fetches a specific contribution
func GetContribution(ctx context.Context, contributionID string) (*UserContribution, error) {
	var contribution UserContribution

	// Make request to User Service
	resp, err := Clients.UserClient.client.Get(ctx, "/api/v1/contributions/"+contributionID, nil)
	if err != nil {
		utils.Logger.Error("Failed to call User Service for getting contribution", zap.Error(err))
		return nil, err
	}

	// Handle response
	if err := handleServiceResponse(resp, &contribution); err != nil {
		return nil, err
	}

	return &contribution, nil
}
