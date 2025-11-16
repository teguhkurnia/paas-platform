package usecase

import (
	"fmt"
	"gofiber-boilerplate/internal/model"
	"gofiber-boilerplate/internal/repository"
	"os"
	"os/exec"

	"github.com/docker/docker/client"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ServiceUseCase struct {
	DB          *gorm.DB
	ServiceRepo *repository.ServiceRepository
	Log         *logrus.Logger
	Validate    *validator.Validate
	Docker      *client.Client
}

func NewServiceUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	serviceRepo *repository.ServiceRepository,
	validate *validator.Validate,
	docker *client.Client,
) *ServiceUseCase {
	return &ServiceUseCase{
		DB:          db,
		Log:         log,
		ServiceRepo: serviceRepo,
		Validate:    validate,
		Docker:      docker,
	}
}

func (u *ServiceUseCase) BuildService(request *model.ServiceBuildRequest) (*string, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		u.Log.Warnf("Invalid request body: %v", err)
		return nil, fiber.ErrBadRequest
	}

	// TODO:
	// Implement service build using service from database
	projectName := "sample-project"
	serviceName := "sample-service"
	clonePath := "./tmp/git/a29ea9e9-6381-4285-b2a4-3b9fd5173a06"
	imageName := slug.Make(fmt.Sprintf("%s-%s", projectName, serviceName))

	u.Log.Debugf("Building service image %s", imageName)

	cmd := exec.Command("railpack", "build", clonePath, "--name", imageName)

	env := os.Environ()
	env = append(env, "NIXPACKS_NODE_VERSION=23")
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		u.Log.Errorf("Failed to build service image: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	u.Log.Infof("Service image %s built successfully", imageName)

	return &imageName, nil
}
