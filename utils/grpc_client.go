package utils

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// GRPCClientManager manages gRPC connections to microservices
type GRPCClientManager struct {
	connections map[string]*grpc.ClientConn
	clients     map[string]interface{}
}

// NewGRPCClientManager creates a new gRPC client manager
func NewGRPCClientManager() *GRPCClientManager {
	return &GRPCClientManager{
		connections: make(map[string]*grpc.ClientConn),
		clients:     make(map[string]interface{}),
	}
}

// Connect establishes a gRPC connection to a service
func (g *GRPCClientManager) Connect(serviceName, address string) (*grpc.ClientConn, error) {
	// Check if connection already exists
	if conn, exists := g.connections[serviceName]; exists {
		return conn, nil
	}

	// Create new connection
	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(30*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s service: %w", serviceName, err)
	}

	// Store connection
	g.connections[serviceName] = conn
	Logger.Info("Connected to gRPC service",
		zap.String("service", serviceName),
		zap.String("address", address),
	)

	return conn, nil
}

// GetConnection returns an existing connection or creates a new one
func (g *GRPCClientManager) GetConnection(serviceName string) (*grpc.ClientConn, error) {
	if conn, exists := g.connections[serviceName]; exists {
		return conn, nil
	}
	return nil, fmt.Errorf("no connection found for service: %s", serviceName)
}

// CloseAll closes all gRPC connections
func (g *GRPCClientManager) CloseAll() {
	for serviceName, conn := range g.connections {
		if err := conn.Close(); err != nil {
			Logger.Error("Failed to close gRPC connection",
				zap.String("service", serviceName),
				zap.Error(err),
			)
		} else {
			Logger.Info("Closed gRPC connection", zap.String("service", serviceName))
		}
	}
	g.connections = make(map[string]*grpc.ClientConn)
	g.clients = make(map[string]interface{})
}

// CreateContextWithRequestID creates a context with request ID for tracing
func CreateContextWithRequestID(ctx context.Context, requestID string) context.Context {
	if requestID != "" {
		md := metadata.New(map[string]string{
			"x-request-id": requestID,
		})
		return metadata.NewOutgoingContext(ctx, md)
	}
	return ctx
}

// Global gRPC client manager instance
var GRPCClients *GRPCClientManager
