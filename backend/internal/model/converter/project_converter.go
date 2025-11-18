package converter

import (
	"gofiber-boilerplate/internal/entity"
	"gofiber-boilerplate/internal/model"
)

func ProjectToResponse(project *entity.Project) *model.ProjectResponse {
	services := make([]model.ServiceResponse, 0)
	for _, service := range project.Services {
		services = append(services, model.ServiceResponse{
			ID:        uint(service.ID),
			Name:      service.Name,
			ProjectID: uint(service.ProjectID),
		})
	}

	return &model.ProjectResponse{
		ID:          uint(project.ID),
		Name:        project.Name,
		Description: project.Description,
		Environment: project.Environment,
		Services:    &services,
	}
}
