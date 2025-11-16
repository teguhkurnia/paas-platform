package usecase

import (
	"context"
	"gofiber-boilerplate/internal/entity"
	"gofiber-boilerplate/internal/model"
	"gofiber-boilerplate/internal/model/converter"
	"gofiber-boilerplate/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProjectUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *repository.ProjectRepository
}

func NewProjectUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	repo *repository.ProjectRepository,
) *ProjectUseCase {
	return &ProjectUseCase{
		DB:         db,
		Log:        log,
		Validate:   validate,
		Repository: repo,
	}
}

func (c *ProjectUseCase) GetAll(
	ctx context.Context,
) (*model.PageResponse[model.ProjectResponse], error) {
	pagination := repository.Pagination[entity.Project]{
		Limit: 50,
		Page:  1,
	}

	projects, err := c.Repository.FindAll(c.DB, pagination)

	c.Log.Debugf("Projects found: %v", projects)

	if err != nil {
		c.Log.Errorf("Failed to get projects: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.ProjectResponse, 0)
	for _, project := range projects.Rows {
		responses = append(responses, *converter.ProjectToResponse(&project))
	}

	return &model.PageResponse[model.ProjectResponse]{
		PageMetaData: model.PageMetaData{
			Page:       projects.Page,
			Size:       projects.TotalPages,
			TotalItems: projects.TotalRows,
			TotalPages: projects.TotalPages,
		},
		Data: responses,
	}, nil
}

func (c *ProjectUseCase) Create(
	ctx context.Context,
	request *model.CreateProjectRequest,
) (*model.ProjectResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body: %v", err)
		return nil, fiber.ErrBadRequest
	}

	project := &entity.Project{
		Name:        request.Name,
		Description: request.Description,
		Environment: request.Environment,
	}

	if err := tx.Create(project).Error; err != nil {
		c.Log.Errorf("Failed to create project: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Errorf("Failed to commit transaction: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ProjectToResponse(project), nil
}
