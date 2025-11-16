package main

import (
	"fmt"
	"gofiber-boilerplate/internal/config"
	"gofiber-boilerplate/internal/util"
	"time"
)

func main() {
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	validate := config.NewValidator()
	app := config.NewFiber(viperConfig)
	db := config.NewDatabase(viperConfig, logger, false)
	redis := config.NewRedis(viperConfig)
	tokenUtil := util.NewTokenUtil(redis, viperConfig, logger)
	rateLimiterUtil := util.NewRateLimiterUtil(redis, 5, time.Duration(1*time.Minute))
	gitUtil := util.NewGitUtil(logger)

	// initialize docker client
	dockerClient := config.NewDockerClient(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		App:             app,
		DB:              db,
		Redis:           redis,
		Log:             logger,
		Validate:        validate,
		Config:          viperConfig,
		TokenUtil:       tokenUtil,
		RateLimiterUtil: rateLimiterUtil,
		GitUtil:         gitUtil,
		Docker:          dockerClient,
	})

	port := viperConfig.GetString("app.port")
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
