package services

import (
	"digimon-story-evolution/utils"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

// ServiceClients holds all microservice clients
type ServiceClients struct {
	DigimonClient *DigimonServiceClient
	UserClient    *UserServiceClient
	AdminClient   *AdminServiceClient
}

// DigimonServiceClient handles communication with Digimon Service
type DigimonServiceClient struct {
	client *utils.HTTPClient
}

// UserServiceClient handles communication with User Service
type UserServiceClient struct {
	client *utils.HTTPClient
}

// AdminServiceClient handles communication with Admin Service
type AdminServiceClient struct {
	client *utils.HTTPClient
}

// NewServiceClients creates new instances of all service clients
func NewServiceClients() *ServiceClients {
	return &ServiceClients{
		DigimonClient: NewDigimonServiceClient(),
		UserClient:    NewUserServiceClient(),
		AdminClient:   NewAdminServiceClient(),
	}
}

// NewDigimonServiceClient creates a new Digimon service client
func NewDigimonServiceClient() *DigimonServiceClient {
	config := utils.GlobalConfig.Services.DigimonService
	client := utils.NewHTTPClient(config.BaseURL, config.Timeout)
	return &DigimonServiceClient{client: client}
}

// NewUserServiceClient creates a new User service client
func NewUserServiceClient() *UserServiceClient {
	config := utils.GlobalConfig.Services.UserService
	client := utils.NewHTTPClient(config.BaseURL, config.Timeout)
	return &UserServiceClient{client: client}
}

// NewAdminServiceClient creates a new Admin service client
func NewAdminServiceClient() *AdminServiceClient {
	config := utils.GlobalConfig.Services.AdminService
	client := utils.NewHTTPClient(config.BaseURL, config.Timeout)
	return &AdminServiceClient{client: client}
}

// handleServiceResponse handles common service response logic
func handleServiceResponse(resp *utils.Response, target interface{}) error {
	if resp.StatusCode >= 400 {
		utils.Logger.Error("Service request failed",
			zap.Int("status_code", resp.StatusCode),
			zap.String("response_body", string(resp.Body)),
		)
		return fmt.Errorf("service request failed with status %d", resp.StatusCode)
	}

	if target != nil {
		if err := json.Unmarshal(resp.Body, target); err != nil {
			utils.Logger.Error("Failed to unmarshal service response",
				zap.Error(err),
				zap.String("response_body", string(resp.Body)),
			)
			return fmt.Errorf("failed to parse service response: %w", err)
		}
	}

	return nil
}

// Global service clients instance
var Clients *ServiceClients
