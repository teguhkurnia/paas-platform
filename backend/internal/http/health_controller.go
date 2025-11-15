package http

import (
	"gofiber-boilerplate/internal/usecase"
	"gofiber-boilerplate/internal/util"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type HealthController struct {
	Log             *logrus.Logger
	UseCase         *usecase.HealthUseCase
	RatelimiterUtil *util.RateLimiterUtil
}

func NewHealthController(
	log *logrus.Logger,
	useCase *usecase.HealthUseCase,
	ratelimiterUtil *util.RateLimiterUtil,
) *HealthController {
	return &HealthController{
		Log:             log,
		UseCase:         useCase,
		RatelimiterUtil: ratelimiterUtil,
	}
}

func (c *HealthController) HealthCheck(ctx *fiber.Ctx) error {
	dockerErr := c.UseCase.CheckDockerHealth(ctx.Context())
	if dockerErr != nil {
		c.Log.Errorf("Docker health check failed: %v", dockerErr)
		ctx.Status(fiber.StatusServiceUnavailable)

		return ctx.JSON(fiber.Map{"status": "unhealthy", "reason": "Docker unhealthy"})
	}

	return ctx.JSON(fiber.Map{"status": "healthy"})
}
