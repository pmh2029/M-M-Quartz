package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	CreatedAt time.Time `gorm:"column:created_at;not null;type:timestamp;default:current_timestamp" mapstructure:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;type:timestamp;default:current_timestamp" mapstructure:"updated_at" json:"updated_at"`
}

type BaseEntityWithDeleted struct {
	BaseEntity
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;index" mapstructure:"deleted_at" json:"deleted_at"`
}
