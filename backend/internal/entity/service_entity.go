package entity

import "time"

type Service struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement;type:bigint signed"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description *string   `gorm:"type:text"`
	OwnerID     int64     `gorm:"type:bigint signed;not null"`
	ProjectID   int64     `gorm:"type:bigint signed;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:milli"`
}

func (Service) TableName() string { return "services" }
