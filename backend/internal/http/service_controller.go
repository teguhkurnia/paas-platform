package http

import (
	"gofiber-boilerplate/internal/model"
	"gofiber-boilerplate/internal/usecase"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ServiceController struct {
	Log     *logrus.Logger
	UseCase *usecase.ServiceUseCase
}

func NewServiceController(log *logrus.Logger, usecase *usecase.ServiceUseCase) *ServiceController {
	return &ServiceController{
		Log:     log,
		UseCase: usecase,
	}
}

func (c *ServiceController) BuildAndDeployService(ctx *fiber.Ctx) error {
	var request model.ServiceBuildRequest

	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	imageName, err := c.UseCase.BuildService(&request)
	if err != nil {
		c.Log.Warnf("Failed to build service: %v", err)
		return err
	}

	// deployResponse, err := c.UseCase.DeployService(&request, *imageName)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":    "success",
		"imageName": imageName,
		// "deploy":    deployResponse,
	})
}

func (c *ServiceController) CreateService(ctx *fiber.Ctx) error {
	var request model.CreateServiceRequest

	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	service, err := c.UseCase.Create(ctx.UserContext(), &request)
	if err != nil {
		c.Log.Warnf("Failed to create service: %v", err)
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"service": service,
	})
}

func (c *ServiceController) UpdateService(ctx *fiber.Ctx) error {
	var request model.UpdateServiceRequest

	serviceID := ctx.Params("id")
	if serviceID == "" {
		c.Log.Warn("Service ID is required")
		return fiber.ErrBadRequest
	}
	serviceIDUint, err := strconv.ParseUint(strings.TrimSpace(serviceID), 10, 64)
	if err != nil {
		c.Log.Warnf("Invalid service ID: %v", err)
		return fiber.ErrBadRequest
	}

	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	service, err := c.UseCase.Update(ctx.UserContext(), uint(serviceIDUint), &request)
	if err != nil {
		c.Log.Warnf("Failed to update service: %v", err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"service": service,
	})
}
