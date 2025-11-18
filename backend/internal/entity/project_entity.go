package entity

import "time"

type Project struct {
	ID          uint64     `json:"id" gorm:"column:id;primaryKey;autoIncrement;type:bigint signed"`
	Name        string     `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Description *string    `json:"description" gorm:"column:description;type:text;null"`
	Environment *string    `json:"environment" gorm:"column:environment;type:text;null"`
	CreatedAt   time.Time  `json:"created_at" gorm:"column:created_at;type:datetime;not null;default:current_timestamp"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"column:updated_at;type:datetime;not null;default:current_timestamp on update current_timestamp"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"column:deleted_at;type:datetime;index"`

	// Relations
	Services []Service `json:"services" gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Project) TableName() string {
	return "projects"
}
