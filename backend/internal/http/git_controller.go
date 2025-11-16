package http

import (
	"gofiber-boilerplate/internal/model"
	"gofiber-boilerplate/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type GitController struct {
	Log     *logrus.Logger
	UseCase *usecase.GitUseCase
}

func NewGitController(log *logrus.Logger, usecase *usecase.GitUseCase) *GitController {
	return &GitController{
		Log:     log,
		UseCase: usecase,
	}
}

func (c *GitController) ConnectGithub(ctx *fiber.Ctx) error {
	err := c.UseCase.ConnectGithub()
	if err != nil {
		c.Log.Errorf("Failed to connect to Github: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to connect to Github"})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Github connected successfully"})
}

func (c *GitController) GitClone(ctx *fiber.Ctx) error {
	req := new(model.CloneGitRequest)
	if err := ctx.BodyParser(req); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}

	err := c.UseCase.CloneRepository(req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to clone repository"})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Repository cloned successfully"})
}
