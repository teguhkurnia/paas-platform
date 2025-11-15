package route

import (
	"gofiber-boilerplate/internal/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App               *fiber.App
	UserController    *http.UserController
	ProjectController *http.ProjectController
	HealthController  *http.HealthController
	AuthMiddleware    fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupAuthRoute() {

	group := c.App.Group("/api/v1")
	group.Use(c.AuthMiddleware)

	group.Get("profile", func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user")
		return ctx.JSON(fiber.Map{
			"status": "success",
			"user":   user,
		})
	})

	group.Post("projects", c.ProjectController.CreateProject)
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/health", c.HealthController.HealthCheck)

	c.App.Post("api/v1/login", c.UserController.Login)
	c.App.Post("api/v1/register", c.UserController.Register)
	c.App.Get("api/v1/verify-email/:code", c.UserController.Verify)
}
