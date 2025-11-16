package http

import (
	"gofiber-boilerplate/internal/model"
	"gofiber-boilerplate/internal/usecase"

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
