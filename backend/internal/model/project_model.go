package model

type ProjectResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Environment *string `json:"environment"`

	Services *[]ServiceResponse `json:"services,omitempty"`
}

type CreateProjectRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=100"`
	Description *string `json:"description"`
	Environment *string `json:"environment"`
}

type UpdateProjectRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=3,max=100"`
	Description *string `json:"description"`
	Environment *string `json:"environment"`
}
