package http

import (
	"gofiber-boilerplate/internal/model"
	"gofiber-boilerplate/internal/usecase"
	"strconv"

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

func (c *ProjectController) GetAllProjects(ctx *fiber.Ctx) error {
	responses, err := c.UseCase.GetAll(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to get projects: %v", err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(responses)
}

func (c *ProjectController) GetProjectByID(ctx *fiber.Ctx) error {
	projectID := ctx.Params("id")
	projectIDUint, err := strconv.ParseUint(projectID, 10, 64)
	if err != nil {
		c.Log.Warnf("Invalid project ID: %v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.GetByID(ctx.UserContext(), uint(projectIDUint))
	if err != nil {
		c.Log.Warnf("Failed to get project by ID: %v", err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
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

func (c *ProjectController) UpdateProject(ctx *fiber.Ctx) error {
	projectID := ctx.Params("id")
	projectIDUint, err := strconv.ParseUint(projectID, 10, 64)
	if err != nil {
		c.Log.Warnf("Invalid project ID: %v", err)
		return fiber.ErrBadRequest
	}

	request := new(model.UpdateProjectRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Update(ctx.UserContext(), uint(projectIDUint), request)
	if err != nil {
		c.Log.Warnf("Failed to update project: %v", err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *ProjectController) DeleteProject(ctx *fiber.Ctx) error {
	projectID := ctx.Params("id")
	projectIDUint, err := strconv.ParseUint(projectID, 10, 64)
	if err != nil {
		c.Log.Warnf("Invalid project ID: %v", err)
		return fiber.ErrBadRequest
	}

	err = c.UseCase.Delete(ctx.UserContext(), uint(projectIDUint))
	if err != nil {
		c.Log.Warnf("Failed to delete project: %v", err)
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
