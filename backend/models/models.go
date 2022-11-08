package models

import (
	"gorm.io/gorm"
	"time"
)

type (
	Model struct {
		ID        *uint           `json:"id,omitempty" gorm:"primarykey"`
		CreatedAt *time.Time      `json:"createdAt,omitempty"`
		UpdatedAt *time.Time      `json:"updatedAt,omitempty"`
		DeletedAt *gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
	}

	Todo struct {
		Model
		Text *string `json:"text,omitempty" gorm:"not null, unique"`
		Done *bool   `json:"done,omitempty" gorm:"not null, default:false"`
	}
)
