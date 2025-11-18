package entity

import "time"

type Service struct {
	ID             uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned"`
	OwnerID        uint64    `gorm:"type:bigint unsigned;not null"`
	ProjectID      uint64    `gorm:"type:bigint unsigned;not null"`
	Name           string    `gorm:"type:varchar(100);not null"`
	Environment    string    `gorm:"type:text"`
	Description    *string   `gorm:"type:text"`
	Provider       string    `gorm:"type:varchar(50);not null"`
	ProviderInputs string    `gorm:"type:json"`
	CreatedAt      time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime:milli"`
}

func (Service) TableName() string { return "services" }
