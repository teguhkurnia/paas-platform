package usecase

import (
	"context"

	"github.com/docker/docker/client"
)

type HealthUseCase struct {
	Docker *client.Client
}

func NewHealthUseCase(docker *client.Client) *HealthUseCase {
	return &HealthUseCase{
		Docker: docker,
	}
}

func (c *HealthUseCase) CheckDockerHealth(context *context.Context) error {
	_, err := c.Docker.Ping()
	return err
}
