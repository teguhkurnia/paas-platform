package config

import (
	"github.com/docker/docker/client"
	"github.com/spf13/viper"
)

func NewDockerClient(config *viper.Viper) *client.Client {
	c, err := client.NewClientWithOpts(client.WithHost(config.GetString("docker.host")), client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	return c
}
