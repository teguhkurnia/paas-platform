package repository

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

type Pagination[T any] struct {
	Limit      int    `json:"limit,omitempty;query:limit"`
	Page       int    `json:"page,omitempty;query:page"`
	Sort       string `json:"sort,omitempty;query:sort"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	Rows       []T    `json:"rows"`
}

func (p *Pagination[T]) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination[T]) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination[T]) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination[T]) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

func paginate[T any](value interface{}, pagination *Pagination[T], db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	// Perbaikan bug kecil: gunakan GetLimit() untuk menghormati default
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) FindAll(
	db *gorm.DB,
	pagination Pagination[T],
) (*Pagination[T], error) {
	entities := []T{}

	if err := db.Scopes(paginate(entities, &pagination, db)).Find(&entities).Error; err != nil {
		return nil, err
	}

	fmt.Printf("Entities found: %v", entities)

	pagination.Rows = entities

	return &pagination, nil
}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}

func (r *Repository[T]) CountById(db *gorm.DB, id uint) (int64, error) {
	var count int64
	err := db.Model(new(T)).Where("id = ?", id).Count(&count).Error
	return count, err
}

func (r *Repository[T]) FindByID(db *gorm.DB, entity *T, id uint) error {
	return db.Where("id = ?", id).Take(entity).Error
}
