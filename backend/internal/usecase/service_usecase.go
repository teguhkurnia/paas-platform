package usecase

import (
	"context"
	"fmt"
	"gofiber-boilerplate/internal/entity"
	"gofiber-boilerplate/internal/model"
	"gofiber-boilerplate/internal/model/converter"
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

func (u *ServiceUseCase) Create(ctx context.Context, request *model.CreateServiceRequest) (*model.ServiceResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := u.Validate.Struct(request)
	if err != nil {
		u.Log.Warnf("Invalid request body: %v", err)
		return nil, fiber.ErrBadRequest
	}

	userID := uint64(1) // TODO: get from auth context
	if request.Environment == nil {
		// default to project environment
		// empty string
		request.Environment = new(string)
	}

	service := &entity.Service{
		OwnerID:        userID,
		ProjectID:      request.ProjectID,
		Name:           request.Name,
		Description:    request.Description,
		Provider:       request.Provider,
		ProviderInputs: request.ProviderInputs,
		Environment:    *request.Environment,
	}

	if err := u.ServiceRepo.Create(tx, service); err != nil {
		u.Log.Errorf("Failed to create service: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.Errorf("Failed to commit transaction: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	response := converter.ServiceToResponse(service)

	return response, nil
}

func (u *ServiceUseCase) Update(ctx context.Context, id uint, request *model.UpdateServiceRequest) (*model.ServiceResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := u.Validate.Struct(request)
	if err != nil {
		u.Log.Warnf("Invalid request body: %v", err)
		return nil, fiber.ErrBadRequest
	}

	service := &entity.Service{}
	err = u.ServiceRepo.FindByID(tx, service, id)
	if err != nil {
		u.Log.Errorf("Failed to get service by ID: %v", err)
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.ErrNotFound
		}
		return nil, fiber.ErrInternalServerError
	}

	if request.Name != nil {
		service.Name = *request.Name
	}
	if request.ProviderInputs != nil {
		service.ProviderInputs = *request.ProviderInputs
	}
	service.Description = request.Description
	service.Environment = func() string {
		if request.Environment != nil {
			return *request.Environment
		}
		return service.Environment
	}()

	if err := u.ServiceRepo.Update(tx, service); err != nil {
		u.Log.Errorf("Failed to update service: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.Errorf("Failed to commit transaction: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	response := converter.ServiceToResponse(service)

	return response, nil
}
