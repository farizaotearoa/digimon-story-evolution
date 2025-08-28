package main

import (
	"context"
	"digimon-story-evolution/routes"
	"digimon-story-evolution/utils"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func setupApplication() *gin.Engine {
	// Initialize configuration from environment variables
	utils.LoadConfig()

	// Initialize Logger
	utils.InitLogger()
	defer utils.Logger.Sync()

	// Connect To PostgreSQL
	utils.ConnectDatabase()

	// Set up routes with custom middleware
	r := gin.New()

	// Add custom middleware
	r.Use(utils.RecoveryMiddleware())
	r.Use(utils.LoggerMiddleware())
	r.Use(utils.RequestIDMiddleware())

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins: utils.GlobalConfig.CORS.AllowOrigins,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "X-Request-ID"},
	}))

	routes.SetupRoutes(r)
	routes.SetupImagesRoutes(r)

	return r
}

func runServer(router *gin.Engine) {
	appName := utils.GlobalConfig.App.Name
	utils.Logger.Info(fmt.Sprintf("Starting %s", appName))

	port := utils.GlobalConfig.App.Port
	if port == "" {
		port = os.Getenv("PORT")
	}

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Start server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	utils.Logger.Info("Shutting down server...")

	// Give outstanding requests additional time for completion
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		utils.Logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	utils.Logger.Info("Server exited")
}

func main() {
	router := setupApplication()
	runServer(router)
}
