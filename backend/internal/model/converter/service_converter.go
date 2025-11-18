package converter

import (
	"gofiber-boilerplate/internal/entity"
	"gofiber-boilerplate/internal/model"
)

func ServiceToResponse(service *entity.Service) *model.ServiceResponse {
	return &model.ServiceResponse{
		ID:        uint(service.ID),
		Name:      service.Name,
		ProjectID: uint(service.ID),
	}
}
