package config

import (
	"gofiber-boilerplate/internal/http"
	"gofiber-boilerplate/internal/http/middleware"
	"gofiber-boilerplate/internal/http/route"
	"gofiber-boilerplate/internal/repository"
	"gofiber-boilerplate/internal/usecase"
	"gofiber-boilerplate/internal/util"

	docker "github.com/docker/docker/client"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App             *fiber.App
	DB              *gorm.DB
	Redis           *redis.Client
	Log             *logrus.Logger
	Validate        *validator.Validate
	Config          *viper.Viper
	TokenUtil       *util.TokenUtil
	RateLimiterUtil *util.RateLimiterUtil
	GitUtil         *util.GitUtil
	Docker          *docker.Client
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepo := repository.NewUserRepository(config.Log)
	projectRepo := repository.NewProjectRepository(config.Log)
	serviceRepo := repository.NewServiceRepository(config.Log)

	// setup usecases
	userUseCase := usecase.NewUserUseCase(
		config.DB,
		config.TokenUtil,
		config.Log,
		config.Validate,
		userRepo,
	)

	healthUseCase := usecase.NewHealthUseCase(
		config.Docker,
	)

	projectUseCase := usecase.NewProjectUseCase(
		config.DB,
		config.Log,
		config.Validate,
		projectRepo,
	)

	gitUseCase := usecase.NewGitUseCase(
		config.Config,
		config.Log,
		config.Validate,
		config.GitUtil,
	)

	serviceUseCase := usecase.NewServiceUseCase(
		config.DB,
		config.Log,
		serviceRepo,
		config.Validate,
		config.Docker,
	)

	// setup controllers
	userController := http.NewUserController(
		config.Log,
		userUseCase,
		config.RateLimiterUtil,
	)

	healthController := http.NewHealthController(
		config.Log,
		healthUseCase,
		config.RateLimiterUtil,
	)

	projectController := http.NewProjectController(
		config.Log,
		projectUseCase,
	)

	gitController := http.NewGitController(
		config.Log,
		gitUseCase,
	)

	serviceController := http.NewServiceController(
		config.Log,
		serviceUseCase,
	)

	// setup routes
	authMiddleware := middleware.NewAuthMiddleware(userUseCase)
	routeConfig := &route.RouteConfig{
		App:               config.App,
		UserController:    userController,
		HealthController:  healthController,
		ProjectController: projectController,
		GitController:     gitController,
		ServiceController: serviceController,
		AuthMiddleware:    authMiddleware,
	}

	routeConfig.Setup()
}
