package services

import (
	"context"
	"digimon-story-evolution/utils"

	"go.uber.org/zap"
)

// PendingContribution represents a contribution waiting for admin approval
type PendingContribution struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	DigimonID   string `json:"digimon_id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UserEmail   string `json:"user_email"`
	UserName    string `json:"user_name"`
}

// ApprovalDecision represents an admin's decision on a contribution
type ApprovalDecision struct {
	ContributionID string `json:"contribution_id"`
	AdminID        string `json:"admin_id"`
	Decision       string `json:"decision"` // "approve", "reject"
	Comments       string `json:"comments"`
}

// GetPendingContributions fetches all pending contributions for admin review
func GetPendingContributions(ctx context.Context) ([]PendingContribution, error) {
	var contributions []PendingContribution

	// Make request to Admin Service
	resp, err := Clients.AdminClient.client.Get(ctx, "/api/v1/admin/contributions/pending", nil)
	if err != nil {
		utils.Logger.Error("Failed to call Admin Service for getting pending contributions", zap.Error(err))
		return nil, err
	}

	// Handle response
	if err := handleServiceResponse(resp, &contributions); err != nil {
		return nil, err
	}

	return contributions, nil
}

// ApproveContribution approves or rejects a contribution
func ApproveContribution(ctx context.Context, decision ApprovalDecision) error {
	// Make request to Admin Service
	resp, err := Clients.AdminClient.client.Post(ctx, "/api/v1/admin/contributions/approve", decision, nil)
	if err != nil {
		utils.Logger.Error("Failed to call Admin Service for approving contribution", zap.Error(err))
		return err
	}

	// Handle response
	if err := handleServiceResponse(resp, nil); err != nil {
		return err
	}

	return nil
}

// GetContributionStats fetches statistics for admin dashboard
func GetContributionStats(ctx context.Context) (map[string]interface{}, error) {
	var stats map[string]interface{}

	// Make request to Admin Service
	resp, err := Clients.AdminClient.client.Get(ctx, "/api/v1/admin/stats", nil)
	if err != nil {
		utils.Logger.Error("Failed to call Admin Service for getting stats", zap.Error(err))
		return nil, err
	}

	// Handle response
	if err := handleServiceResponse(resp, &stats); err != nil {
		return nil, err
	}

	return stats, nil
}
