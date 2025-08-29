package services

import (
	"context"
	"digimon-story-evolution/proto"
	"digimon-story-evolution/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// GRPCDigimonServiceClient handles gRPC communication with Digimon Service
type GRPCDigimonServiceClient struct {
	client proto.DigimonServiceClient
	conn   *grpc.ClientConn
}

// GRPCUserServiceClient handles gRPC communication with User Service
type GRPCUserServiceClient struct {
	client proto.UserServiceClient
	conn   *grpc.ClientConn
}

// GRPCAdminServiceClient handles gRPC communication with Admin Service
type GRPCAdminServiceClient struct {
	client proto.AdminServiceClient
	conn   *grpc.ClientConn
}

// NewGRPCDigimonServiceClient creates a new gRPC Digimon service client
func NewGRPCDigimonServiceClient() (*GRPCDigimonServiceClient, error) {
	conn, err := utils.GRPCClients.Connect("digimon", utils.GlobalConfig.Services.DigimonService.GRPCAddress)
	if err != nil {
		return nil, err
	}

	client := proto.NewDigimonServiceClient(conn)
	return &GRPCDigimonServiceClient{
		client: client,
		conn:   conn,
	}, nil
}

// NewGRPCUserServiceClient creates a new gRPC User service client
func NewGRPCUserServiceClient() (*GRPCUserServiceClient, error) {
	conn, err := utils.GRPCClients.Connect("user", utils.GlobalConfig.Services.UserService.GRPCAddress)
	if err != nil {
		return nil, err
	}

	client := proto.NewUserServiceClient(conn)
	return &GRPCUserServiceClient{
		client: client,
		conn:   conn,
	}, nil
}

// NewGRPCAdminServiceClient creates a new gRPC Admin service client
func NewGRPCAdminServiceClient() (*GRPCAdminServiceClient, error) {
	conn, err := utils.GRPCClients.Connect("admin", utils.GlobalConfig.Services.AdminService.GRPCAddress)
	if err != nil {
		return nil, err
	}

	client := proto.NewAdminServiceClient(conn)
	return &GRPCAdminServiceClient{
		client: client,
		conn:   conn,
	}, nil
}

// GetDigimonList calls the Digimon service via gRPC
func (d *GRPCDigimonServiceClient) GetDigimonList(ctx context.Context, req *proto.GetDigimonListRequest) (*proto.GetDigimonListResponse, error) {
	utils.Logger.Debug("Calling Digimon service GetDigimonList via gRPC")

	response, err := d.client.GetDigimonList(ctx, req)
	if err != nil {
		utils.Logger.Error("Failed to call Digimon service GetDigimonList", zap.Error(err))
		return nil, err
	}

	return response, nil
}

// GetDigimonListSize calls the Digimon service via gRPC
func (d *GRPCDigimonServiceClient) GetDigimonListSize(ctx context.Context, req *proto.GetDigimonListRequest) (*proto.GetDigimonListSizeResponse, error) {
	utils.Logger.Debug("Calling Digimon service GetDigimonListSize via gRPC")

	response, err := d.client.GetDigimonListSize(ctx, req)
	if err != nil {
		utils.Logger.Error("Failed to call Digimon service GetDigimonListSize", zap.Error(err))
		return nil, err
	}

	return response, nil
}

// GetDigimonDetails calls the Digimon service via gRPC
func (d *GRPCDigimonServiceClient) GetDigimonDetails(ctx context.Context, req *proto.GetDigimonDetailsRequest) (*proto.GetDigimonDetailsResponse, error) {
	utils.Logger.Debug("Calling Digimon service GetDigimonDetails via gRPC")

	response, err := d.client.GetDigimonDetails(ctx, req)
	if err != nil {
		utils.Logger.Error("Failed to call Digimon service GetDigimonDetails", zap.Error(err))
		return nil, err
	}

	return response, nil
}

// GetDigimonEvolutions calls the Digimon service via gRPC
func (d *GRPCDigimonServiceClient) GetDigimonEvolutions(ctx context.Context, req *proto.GetDigimonDetailsRequest) (*proto.GetDigimonEvolutionsResponse, error) {
	utils.Logger.Debug("Calling Digimon service GetDigimonEvolutions via gRPC")

	response, err := d.client.GetDigimonEvolutions(ctx, req)
	if err != nil {
		utils.Logger.Error("Failed to call Digimon service GetDigimonEvolutions", zap.Error(err))
		return nil, err
	}

	return response, nil
}

// CreateContribution calls the User service via gRPC
func (u *GRPCUserServiceClient) CreateContribution(ctx context.Context, req *proto.CreateContributionRequest) (*proto.CreateContributionResponse, error) {
	utils.Logger.Debug("Calling User service CreateContribution via gRPC")

	response, err := u.client.CreateContribution(ctx, req)
	if err != nil {
		utils.Logger.Error("Failed to call User service CreateContribution", zap.Error(err))
		return nil, err
	}

	return response, nil
}

// GetUserContributions calls the User service via gRPC
func (u *GRPCUserServiceClient) GetUserContributions(ctx context.Context, req *proto.GetUserContributionsRequest) (*proto.GetUserContributionsResponse, error) {
	utils.Logger.Debug("Calling User service GetUserContributions via gRPC")

	response, err := u.client.GetUserContributions(ctx, req)
	if err != nil {
		utils.Logger.Error("Failed to call User service GetUserContributions", zap.Error(err))
		return nil, err
	}

	return response, nil
}

// GetContribution calls the User service via gRPC
func (u *GRPCUserServiceClient) GetContribution(ctx context.Context, req *proto.GetContributionRequest) (*proto.GetContributionResponse, error) {
	utils.Logger.Debug("Calling User service GetContribution via gRPC")

	response, err := u.client.GetContribution(ctx, req)
	if err != nil {
		utils.Logger.Error("Failed to call User service GetContribution", zap.Error(err))
		return nil, err
	}

	return response, nil
}

// GetPendingContributions calls the Admin service via gRPC
func (a *GRPCAdminServiceClient) GetPendingContributions(ctx context.Context, req *proto.GetPendingContributionsRequest) (*proto.GetPendingContributionsResponse, error) {
	utils.Logger.Debug("Calling Admin service GetPendingContributions via gRPC")

	response, err := a.client.GetPendingContributions(ctx, req)
	if err != nil {
		utils.Logger.Error("Failed to call Admin service GetPendingContributions", zap.Error(err))
		return nil, err
	}

	return response, nil
}

// ApproveContribution calls the Admin service via gRPC
func (a *GRPCAdminServiceClient) ApproveContribution(ctx context.Context, req *proto.ApproveContributionRequest) (*proto.ApproveContributionResponse, error) {
	utils.Logger.Debug("Calling Admin service ApproveContribution via gRPC")

	response, err := a.client.ApproveContribution(ctx, req)
	if err != nil {
		utils.Logger.Error("Failed to call Admin service ApproveContribution", zap.Error(err))
		return nil, err
	}

	return response, nil
}

// GetContributionStats calls the Admin service via gRPC
func (a *GRPCAdminServiceClient) GetContributionStats(ctx context.Context, req *proto.GetContributionStatsRequest) (*proto.GetContributionStatsResponse, error) {
	utils.Logger.Debug("Calling Admin service GetContributionStats via gRPC")

	response, err := a.client.GetContributionStats(ctx, req)
	if err != nil {
		utils.Logger.Error("Failed to call Admin service GetContributionStats", zap.Error(err))
		return nil, err
	}

	return response, nil
}

// GetDashboardOverview calls the Admin service via gRPC
func (a *GRPCAdminServiceClient) GetDashboardOverview(ctx context.Context, req *proto.GetDashboardOverviewRequest) (*proto.GetDashboardOverviewResponse, error) {
	utils.Logger.Debug("Calling Admin service GetDashboardOverview via gRPC")

	response, err := a.client.GetDashboardOverview(ctx, req)
	if err != nil {
		utils.Logger.Error("Failed to call Admin service GetDashboardOverview", zap.Error(err))
		return nil, err
	}

	return response, nil
}
