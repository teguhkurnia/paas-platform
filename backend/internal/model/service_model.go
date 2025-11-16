package model

type ServiceBuildRequest struct {
	ServiceID uint `json:"service_id" validate:"required"`
}
