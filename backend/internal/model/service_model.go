package model

type ServiceBuildRequest struct {
	ServiceID uint `json:"service_id" validate:"required"`
}

type ServiceResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ProjectID uint   `json:"project_id"`
}

type CreateServiceRequest struct {
	ProjectID      uint64  `json:"project_id" validate:"required"`
	Name           string  `json:"name" validate:"required,min=3,max=100"`
	Description    *string `json:"description"`
	Provider       string  `json:"provider" validate:"required"`
	ProviderInputs string  `json:"provider_inputs" validate:"required,json"`
	Environment    *string `json:"environment"`
}

type UpdateServiceRequest struct {
	Name           *string `json:"name" validate:"omitempty,min=3,max=100"`
	Description    *string `json:"description"`
	ProviderInputs *string `json:"provider_inputs" validate:"omitempty,json"`
	Environment    *string `json:"environment"`
}
