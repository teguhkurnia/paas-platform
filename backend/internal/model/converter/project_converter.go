package converter

import (
	"gofiber-boilerplate/internal/entity"
	"gofiber-boilerplate/internal/model"
)

func ProjectToResponse(project *entity.Project) *model.ProjectResponse {
	return &model.ProjectResponse{
		ID:          uint(project.ID),
		Name:        project.Name,
		Description: project.Description,
	}
}
