package entity

import "time"

type Project struct {
	ID          int64      `json:"id" gorm:"column:id;primaryKey;autoIncrement;type:bigint signed"`
	Name        string     `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Description *string    `json:"description" gorm:"column:description;type:text;null"`
	CreatedAt   time.Time  `json:"created_at" gorm:"column:created_at;type:datetime;not null;default:current_timestamp"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"column:updated_at;type:datetime;not null;default:current_timestamp on update current_timestamp"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"column:deleted_at;type:datetime;index"`
}

func (Project) TableName() string {
	return "projects"
}
