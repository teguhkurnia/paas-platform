package http

import (
	"gofiber-boilerplate/internal/model"
	"gofiber-boilerplate/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProjectController struct {
	Log     *logrus.Logger
	UseCase *usecase.ProjectUseCase
}

func NewProjectController(
	log *logrus.Logger,
	useCase *usecase.ProjectUseCase,
) *ProjectController {
	return &ProjectController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *ProjectController) CreateProject(ctx *fiber.Ctx) error {
	request := new(model.CreateProjectRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create project: %v", err)
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
